package office

import (
	"context"
	"example/app/company/cmd/rpc/office/pb"

	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOfficeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOfficeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOfficeLogic {
	return &GetOfficeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOfficeLogic) GetOffice(req *types.OfficeDetailRequest) (types.OfficeDetailResponse, error) {
	var resp types.OfficeDetailResponse

	officeResp, err := l.svcCtx.OfficeRpcClient.GetOffice(l.ctx, &pb.GetOfficeRequest{
		Id: req.ID,
	})

	resp.Office = types.Office{
		Name:    officeResp.Office.Name,
		Address: officeResp.Office.Address,
	}

	return resp, err
}
