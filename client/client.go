package main

import (
	"context"
	"log"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/wekeeroad/GoRPC/pkg/middleware"
	pb "github.com/wekeeroad/GoRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

var opts []grpc.DialOption

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
		),
	))
	opts = append(opts, grpc.WithStreamInterceptor(
		grpc_middleware.ChainStreamClient(
			middleware.StreamContextTimeout(),
		),
	))
}

func main() {
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
