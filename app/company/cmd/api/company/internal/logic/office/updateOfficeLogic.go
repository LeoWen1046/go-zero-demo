package office

import (
	"context"
	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"
	"example/app/company/cmd/rpc/office/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOfficeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOfficeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOfficeLogic {
	return &UpdateOfficeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOfficeLogic) UpdateOffice(req *types.OfficeUpdateRequest) (types.OfficeUpdateResponse, error) {
	var resp types.OfficeUpdateResponse

	officeResp, err := l.svcCtx.OfficeRpcClient.UpdateOffice(l.ctx, &pb.UpdateOfficeRequest{
		Id:      req.ID,
		Name:    req.Name,
		Address: req.Address,
	})
	if err != nil {
		resp.Msg = "更新失败"
		return resp, err
	}

	resp.Msg = officeResp.Msg

	return resp, nil
}
