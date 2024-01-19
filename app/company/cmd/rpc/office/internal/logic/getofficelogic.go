package logic

import (
	"context"
	repo "example/app/company/cmd/rpc/office/internal/repository"

	"example/app/company/cmd/rpc/office/internal/svc"
	"example/app/company/cmd/rpc/office/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOfficeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOfficeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOfficeLogic {
	return &GetOfficeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOfficeLogic) GetOffice(in *pb.GetOfficeRequest) (*pb.GetOfficeResponse, error) {
	resp := new(pb.GetOfficeResponse)
	office, getErr := repo.NewOfficeRepo(l.ctx, l.svcCtx).GetOffice(l.ctx, in.Id)
	if getErr != nil {
		return resp, getErr
	}

	resp.Office = &pb.Office{
		Name:    office.Name,
		Address: office.Address,
	}

	return resp, nil
}
