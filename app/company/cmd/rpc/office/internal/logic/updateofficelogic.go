package logic

import (
	"context"
	repo "example/app/company/cmd/rpc/office/internal/repository"
	"example/app/company/model"

	"example/app/company/cmd/rpc/office/internal/svc"
	"example/app/company/cmd/rpc/office/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOfficeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOfficeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOfficeLogic {
	return &UpdateOfficeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOfficeLogic) UpdateOffice(in *pb.UpdateOfficeRequest) (*pb.UpdateOfficeResponse, error) {
	var office model.Office
	resp := new(pb.UpdateOfficeResponse)
	resp.Msg = "更新成功"

	if in.Name != "" {
		office.Name = in.Name
	}
	if in.Address != "" {
		office.Address = in.Address
	}

	office.ID = in.Id

	err := repo.NewOfficeRepo(l.ctx, l.svcCtx).UpdateOffice(l.ctx, office)
	if err != nil {
		resp.Msg = "更新失败"
		return resp, err
	}

	return resp, nil
}
