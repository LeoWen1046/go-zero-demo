package logic

import (
	"context"

	repo "example/app/company/cmd/rpc/office/internal/repository"
	"example/app/company/cmd/rpc/office/internal/svc"
	"example/app/company/cmd/rpc/office/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListOfficeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListOfficeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOfficeLogic {
	return &ListOfficeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListOfficeLogic) ListOffice(in *pb.ListOfficeRequest) (*pb.ListOfficeResponse, error) {
	resp := new(pb.ListOfficeResponse)
	cnt, offices, err := repo.NewOfficeRepo(l.ctx, l.svcCtx).ListOffice(l.ctx, *in)
	if err != nil {
		return resp, err
	}
	resp.Total = cnt
	for _, office := range offices {
		resp.List = append(resp.List, &pb.Office{
			Id:      office.BaseModel.ID,
			Name:    office.Name,
			Address: office.Address,
		})
	}

	return resp, nil
}
