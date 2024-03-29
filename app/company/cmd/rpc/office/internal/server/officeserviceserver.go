// Code generated by goctl. DO NOT EDIT.
// Source: office.proto

package server

import (
	"context"

	"example/app/company/cmd/rpc/office/internal/logic"
	"example/app/company/cmd/rpc/office/internal/svc"
	"example/app/company/cmd/rpc/office/pb"
)

type OfficeServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOfficeServiceServer
}

func NewOfficeServiceServer(svcCtx *svc.ServiceContext) *OfficeServiceServer {
	return &OfficeServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *OfficeServiceServer) CreatOffice(ctx context.Context, in *pb.CreateOfficeRequest) (*pb.CreateOfficeResponse, error) {
	l := logic.NewCreatOfficeLogic(ctx, s.svcCtx)
	return l.CreatOffice(in)
}

func (s *OfficeServiceServer) DeleteOffice(ctx context.Context, in *pb.DeleteOfficeRequest) (*pb.DeleteOfficeResponse, error) {
	l := logic.NewDeleteOfficeLogic(ctx, s.svcCtx)
	return l.DeleteOffice(in)
}

func (s *OfficeServiceServer) UpdateOffice(ctx context.Context, in *pb.UpdateOfficeRequest) (*pb.UpdateOfficeResponse, error) {
	l := logic.NewUpdateOfficeLogic(ctx, s.svcCtx)
	return l.UpdateOffice(in)
}

func (s *OfficeServiceServer) GetOffice(ctx context.Context, in *pb.GetOfficeRequest) (*pb.GetOfficeResponse, error) {
	l := logic.NewGetOfficeLogic(ctx, s.svcCtx)
	return l.GetOffice(in)
}

func (s *OfficeServiceServer) ListOffice(ctx context.Context, in *pb.ListOfficeRequest) (*pb.ListOfficeResponse, error) {
	l := logic.NewListOfficeLogic(ctx, s.svcCtx)
	return l.ListOffice(in)
}
