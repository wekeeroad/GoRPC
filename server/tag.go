package server

import (
	"context"
	"encoding/json"

	"github.com/wekeeroad/GoRPC/pkg/bapi"
	"github.com/wekeeroad/GoRPC/pkg/errcode"
	pb "github.com/wekeeroad/GoRPC/proto"
	"google.golang.org/grpc/metadata"
)

type TagServer struct {
	auth *Auth
}

type Auth struct{}

func (a *Auth) GetAppKey() string {
	return "go-programming-tour-book"
}

func (a *Auth) GetAppSecret() string {
	return "eddycjy"
}

func (a *Auth) Check(ctx context.Context) error {
	md, _ := metadata.FromIncomingContext(ctx)

	var appKey, appSecret string
	if value, ok := md["app_key"]; ok {
		appKey = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}
	if appKey != a.GetAppKey() || appSecret != a.GetAppSecret() {
		return errcode.ToRPCError(errcode.Unauthorized)
	}
	return nil
}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	if err := t.auth.Check(ctx); err != nil {
		return nil, err
	}
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, errcode.ToRPCError(errcode.ErrorGetTagListFail)
	}
	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		//internal_err := errcode.ToRPCError(errcode.InvalidParams)
		//sts := errcode.FormError(internal_err)
		//fmt.Printf("The detail of err: %q\n", sts.Details())
		return nil, errcode.ToRPCError(errcode.InvalidParams)
	}
	return &tagList, nil
}
