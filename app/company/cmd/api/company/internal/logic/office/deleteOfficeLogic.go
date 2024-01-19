package office

import (
	"context"
	"example/app/company/cmd/rpc/office/pb"

	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOfficeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteOfficeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOfficeLogic {
	return &DeleteOfficeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteOfficeLogic) DeleteOffice(req *types.OfficeDeleteRequest) (types.OfficeDeleteResponse, error) {
	var resp types.OfficeDeleteResponse

	officeResp, err := l.svcCtx.OfficeRpcClient.DeleteOffice(l.ctx, &pb.DeleteOfficeRequest{
		Id: req.ID,
	})

	resp.Msg = officeResp.Msg

	return resp, err
}
