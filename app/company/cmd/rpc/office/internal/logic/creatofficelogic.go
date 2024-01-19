package logic

import (
	"context"
	repo "example/app/company/cmd/rpc/office/internal/repository"
	"example/app/company/model"

	"example/app/company/cmd/rpc/office/internal/svc"
	"example/app/company/cmd/rpc/office/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatOfficeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatOfficeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatOfficeLogic {
	return &CreatOfficeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatOfficeLogic) CreatOffice(in *pb.CreateOfficeRequest) (*pb.CreateOfficeResponse, error) {
	office := model.Office{
		Name:    in.Name,
		Address: in.Address,
	}
	resp := new(pb.CreateOfficeResponse)
	resp.Msg = "创建成功"

	err := repo.NewOfficeRepo(l.ctx, l.svcCtx).CreateOffice(l.ctx, office)
	if err != nil {
		resp.Msg = "创建失败"
		return resp, err
	}

	return resp, nil
}
