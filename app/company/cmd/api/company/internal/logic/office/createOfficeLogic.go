package office

import (
	"context"
	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"
	"example/app/company/cmd/rpc/office/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOfficeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOfficeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOfficeLogic {
	return &CreateOfficeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOfficeLogic) CreateOffice(req *types.OfficeCreateRequest) (types.OfficeCreateResponse, error) {
	var resp types.OfficeCreateResponse

	officeResp, err := l.svcCtx.OfficeRpcClient.CreatOffice(l.ctx, &pb.CreateOfficeRequest{
		Name:    req.Name,
		Address: req.Address,
	})

	resp.Msg = officeResp.Msg

	return resp, err
}
