package employee

import (
	"context"
	"example/app/company/model"

	repo "example/app/company/cmd/api/company/internal/repository"
	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateEmployeeLogic {
	return &CreateEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateEmployeeLogic) CreateEmployee(req *types.EmployeeCreateRequest) (types.EmployeeCreateResponse, error) {
	emp := model.Employee{
		Name:     req.Name,
		Age:      req.Age,
		Company:  req.Company,
		Phone:    req.Phone,
		OfficeID: req.OfficeID,
	}
	var resp types.EmployeeCreateResponse
	resp.Msg = "创建成功"

	err := repo.NewEmployeeRepo(l.ctx, l.svcCtx).CreateEmployee(l.ctx, emp)
	if err != nil {
		resp.Msg = "创建失败"
		return resp, err
	}

	return resp, nil
}
