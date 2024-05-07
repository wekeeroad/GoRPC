package server

import (
	"context"
	"encoding/json"

	"github.com/wekeeroad/GoRPC/pkg/bapi"
	"github.com/wekeeroad/GoRPC/pkg/errcode"
	pb "github.com/wekeeroad/GoRPC/proto"
)

type TagServer struct{}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
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
