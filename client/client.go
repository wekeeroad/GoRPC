package main

import (
	"context"
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/wekeeroad/GoRPC/global"
	"github.com/wekeeroad/GoRPC/pkg/middleware"
	"github.com/wekeeroad/GoRPC/pkg/tracer"
	pb "github.com/wekeeroad/GoRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var opts []grpc.DialOption

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
	opts = append(opts, grpc.WithUnaryInterceptor(
		grpc_middleware.ChainUnaryClient(
			grpc_retry.UnaryClientInterceptor( // 重试的中间件要注册在timeout中间件之前，不然不生效。
				grpc_retry.WithMax(2),
				grpc_retry.WithCodes(
					codes.Unknown,
					codes.Internal,
					codes.DeadlineExceeded,
				),
			),
			middleware.UnaryContextTimeout(),
			//middleware.ClientTracing(),
		),
	))
	opts = append(opts, grpc.WithStreamInterceptor(
		grpc_middleware.ChainStreamClient(
			middleware.StreamContextTimeout(),
		),
	))
	/*
		err := setupTracer()
		if err != nil {
			log.Fatalf("init.setupTracer err: %v", err)
		}
	*/
}

func main() {
	auth := Auth{
		AppKey:    "go-programming-tour-book",
		AppSecret: "eddycjy",
	}
	opts = append(opts, grpc.WithPerRPCCredentials(&auth))
	ctx := context.Background()
	clientConn, _ := GetClientConn(ctx, "localhost:9003", opts)
	defer clientConn.Close()
	tagServiceClient := pb.NewTagServiceClient(clientConn)
	resp, _ := tagServiceClient.GetTagList(
		ctx,
		&pb.GetTagListRequest{Name: "luk"},
	)
	log.Printf("resp: %v", resp)
}

func GetClientConn(ctx context.Context, target string, opts []grpc.DialOption) (*grpc.ClientConn, error) {
	opts = append(opts, grpc.WithInsecure())
	return grpc.DialContext(ctx, target, opts...)
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer(
		"tag-client",
		"127.0.0.1:6831",
	)
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
