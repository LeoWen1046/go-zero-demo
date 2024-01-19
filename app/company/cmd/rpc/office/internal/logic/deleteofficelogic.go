package logic

import (
	"context"
	repo "example/app/company/cmd/rpc/office/internal/repository"

	"example/app/company/cmd/rpc/office/internal/svc"
	"example/app/company/cmd/rpc/office/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOfficeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOfficeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOfficeLogic {
	return &DeleteOfficeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOfficeLogic) DeleteOffice(in *pb.DeleteOfficeRequest) (*pb.DeleteOfficeResponse, error) {
	resp := new(pb.DeleteOfficeResponse)
	resp.Msg = "删除成功"

	err := repo.NewOfficeRepo(l.ctx, l.svcCtx).DeleteOffice(l.ctx, in.Id)
	if err != nil {
		resp.Msg = "删除失败"
		return resp, err
	}
	return resp, nil
}
