package office

import (
	"context"
	"example/app/company/cmd/rpc/office/pb"

	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListOfficeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListOfficeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOfficeLogic {
	return &ListOfficeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListOfficeLogic) ListOffice(req types.OfficeListRequest) (types.OfficeListResponse, error) {
	var resp types.OfficeListResponse
	resp.List = make([]types.Office, 0)
	offices, err := l.svcCtx.OfficeRpcClient.ListOffice(l.ctx, &pb.ListOfficeRequest{
		Page:     int64(req.Page),
		PageSize: int64(req.PageSize),
		Search:   req.Search,
		Order:    req.Order,
		Sort:     req.Sort,
	})
	if err != nil {
		return resp, err
	}
	resp.Total = int(offices.Total)
	for _, office := range offices.List {
		resp.List = append(resp.List, types.Office{
			ID:      office.Id,
			Name:    office.Name,
			Address: office.Address,
		})
	}

	return resp, nil
}
