package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"net"
	"net/http"
	"path"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/wekeeroad/GoRPC/global"
	"github.com/wekeeroad/GoRPC/pkg/middleware"
	"github.com/wekeeroad/GoRPC/pkg/swagger"
	"github.com/wekeeroad/GoRPC/pkg/tracer"
	pb "github.com/wekeeroad/GoRPC/proto"
	"github.com/wekeeroad/GoRPC/server"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var grpcPort string
var port string

type httpError struct {
	Code    int32  `json:"code, omitempty"`
	Message string `json:"message, omitempty"`
}

type Auth struct {
	AppKey    string
	AppSecret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_key": a.AppKey, "app_secret": a.AppSecret}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return false
}

func init() {
	err := setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err: %v", err)
	}
	flag.StringVar(&grpcPort, "grpcPort", "9000", "The port of serve for gRPC")
	flag.StringVar(&port, "port", "9003", "The port of serve for multiRoute")
	flag.Parse()
}

func main() {
	/*
		go func() {
			err := RunGrpcServerDre(grpcPort)
			if err != nil {
				log.Fatalf("Run RunGrpcServerDre err: %v", err)
			}
		}()
	*/
	err := RunServer(port)
	if err != nil {
		log.Fatalf("Run Serve err: %v", err)
	}
}

func RunHttpServer(port string) *http.Server {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	return &http.Server{
		Addr:    ":" + port,
		Handler: serveMux,
	}

}

func RunGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	return s
}

func RunTCPServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func RunGrpcServerDre(port string) error {
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func RunServer(port string) error {
	httpMux := runHttpServer()
	grpcS := runGrpcServer()
	gatewatMux := runGrpcGatwayServer()

	httpMux.Handle("/", gatewatMux)
	return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	prefix := "/swagger-ui/"
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger_ui",
	})
	serveMux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	serveMux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			http.NotFound(w, r)
			return
		}
		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join("proto", p)
		http.ServeFile(w, r, p)
	})
	return serveMux
}

func runGrpcServer() *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			middleware.AccessLog,
			middleware.ErrorLog,
			middleware.Recovery,
			middleware.ContextTimeout,
			middleware.ServerTracing,
		)),
	}
	s := grpc.NewServer(opts...)
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)

	return s
}

func runGrpcGatwayServer() *runtime.ServeMux {
	endpoint := "127.0.0.1:" + port
	runtime.HTTPError = grpcGatewayError
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	// 这一部分参数需要从http1.1的request中获取，并写入到dopts中，不能如下进行硬编码。
	auth := Auth{
		AppKey:    "go-programming-tour-book",
		AppSecret: "eddycjy",
	}
	dopts = append(dopts, grpc.WithPerRPCCredentials(&auth))
	_ = pb.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)

	return gwmux
}

func grpcGatewayError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	httpError := httpError{Code: int32(s.Code()), Message: s.Message()}
	details := s.Details()
	for _, detail := range details {
		if v, ok := detail.(*pb.Error); ok {
			httpError.Code = v.Code
			httpError.Message = v.Message
		}
	}
	resp, _ := json.Marshal(httpError)
	w.Header().Set("Content-type", marshaler.ContentType())
	w.WriteHeader(runtime.HTTPStatusFromCode(s.Code()))
	_, _ = w.Write(resp)
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"tag-service",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
