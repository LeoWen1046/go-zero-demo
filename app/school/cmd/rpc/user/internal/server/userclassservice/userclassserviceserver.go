// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"example/app/school/cmd/rpc/user/github.com/example/user"
	"example/app/school/cmd/rpc/user/internal/logic/userclassservice"
	"example/app/school/cmd/rpc/user/internal/svc"
)

type UserClassServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserClassServiceServer
}

func NewUserClassServiceServer(svcCtx *svc.ServiceContext) *UserClassServiceServer {
	return &UserClassServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserClassServiceServer) UserClassList(ctx context.Context, in *user.UserClassListReq) (*user.UserClassListResp, error) {
	l := userclassservicelogic.NewUserClassListLogic(ctx, s.svcCtx)
	return l.UserClassList(in)
}

func (s *UserClassServiceServer) UserClassUpdate(ctx context.Context, in *user.UserClassUpdateReq) (*user.UserClassUpdateResp, error) {
	l := userclassservicelogic.NewUserClassUpdateLogic(ctx, s.svcCtx)
	return l.UserClassUpdate(in)
}

func (s *UserClassServiceServer) UserClassInfo(ctx context.Context, in *user.UserClassInfoReq) (*user.UserClassInfoResp, error) {
	l := userclassservicelogic.NewUserClassInfoLogic(ctx, s.svcCtx)
	return l.UserClassInfo(in)
}

func (s *UserClassServiceServer) UserClassAdd(ctx context.Context, in *user.UserClassAddReq) (*user.UserClassAddResp, error) {
	l := userclassservicelogic.NewUserClassAddLogic(ctx, s.svcCtx)
	return l.UserClassAdd(in)
}

func (s *UserClassServiceServer) UserClassDelete(ctx context.Context, in *user.UserClassDeleteReq) (*user.UserClassDeleteResp, error) {
	l := userclassservicelogic.NewUserClassDeleteLogic(ctx, s.svcCtx)
	return l.UserClassDelete(in)
}