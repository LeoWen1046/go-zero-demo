package employee

import (
	"context"
	"example/app/company/model"

	repo "example/app/company/cmd/api/company/internal/repository"
	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateEmployeeLogic {
	return &UpdateEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateEmployeeLogic) UpdateEmployee(req *types.EmployeeUpdateRequest) (types.EmployeeUpdateResponse, error) {
	var emp model.Employee
	var resp types.EmployeeUpdateResponse
	resp.Msg = "更新成功"

	if req.Name != "" {
		emp.Name = req.Name
	}
	if req.Age > 0 {
		emp.Age = req.Age
	}
	if req.Company != "" {
		emp.Company = req.Company
	}
	if req.Phone != "" {
		emp.Phone = req.Phone
	}
	if req.OfficeID > 0 {
		emp.OfficeID = req.OfficeID
	}

	emp.ID = req.Id

	err := repo.NewEmployeeRepo(l.ctx, l.svcCtx).UpdateEmployee(l.ctx, emp)
	if err != nil {
		resp.Msg = "更新失败"
		return resp, err
	}

	return resp, nil
}
