package employee

import (
	"context"

	repo "example/app/company/cmd/api/company/internal/repository"
	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListEmployeeLogic {
	return &ListEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListEmployeeLogic) ListEmployee(req types.EmployeeListRequest) (types.EmployeeListResponse, error) {
	var resp types.EmployeeListResponse
	resp.List = make([]types.Employee, 0)
	cnt, emps, err := repo.NewEmployeeRepo(l.ctx, l.svcCtx).ListEmployee(l.ctx, req)
	if err != nil {
		return resp, err
	}
	resp.Total = cnt
	for _, emp := range emps {
		resp.List = append(resp.List, types.Employee{
			Id:       emp.BaseModel.ID,
			Name:     emp.Name,
			Age:      emp.Age,
			Company:  emp.Company,
			Phone:    emp.Phone,
			OfficeID: emp.OfficeID,
		})
	}

	return resp, nil
}
