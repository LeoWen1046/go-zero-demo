package employee

import (
	"context"
	repo "example/app/company/cmd/api/company/internal/repository"

	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmployeeLogic {
	return &GetEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetEmployeeLogic) GetEmployee(req *types.EmployeeDetailRequest) (types.EmployeeDetailResponse, error) {
	var resp types.EmployeeDetailResponse
	emp, getErr := repo.NewEmployeeRepo(l.ctx, l.svcCtx).GetEmployee(l.ctx, req.Id)
	if getErr != nil {
		return resp, getErr
	}

	resp.Employee = types.Employee{
		Id:       emp.BaseModel.ID,
		Name:     emp.Name,
		Age:      emp.Age,
		Company:  emp.Company,
		Phone:    emp.Phone,
		OfficeID: emp.OfficeID,
	}

	return resp, nil
}
