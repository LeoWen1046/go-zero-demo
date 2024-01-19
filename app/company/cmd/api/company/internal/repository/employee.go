package repository

import (
	"context"
	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"
	"example/app/company/model"
)

type EmployeeRepo interface {
	CreateEmployee(ctx context.Context, emp model.Employee) error
	DeleteEmployee(ctx context.Context, Id int64) error
	UpdateEmployee(ctx context.Context, emp model.Employee) error
	GetEmployee(ctx context.Context, Id int64) (model.Employee, error)
	ListEmployee(ctx context.Context, req types.EmployeeListRequest) ([]*model.Employee, error)
}

type employeeRepo struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmployeeRepo(c context.Context, serviceContext *svc.ServiceContext) *employeeRepo {
	return &employeeRepo{
		ctx:    c,
		svcCtx: serviceContext,
	}
}

func (e *employeeRepo) CreateEmployee(ctx context.Context, emp model.Employee) error {
	return e.svcCtx.DbEngin.WithContext(ctx).Create(&emp).Error
}

func (e *employeeRepo) DeleteEmployee(ctx context.Context, Id int64) error {
	return e.svcCtx.DbEngin.WithContext(ctx).Where("id = ?", Id).Delete(&model.Employee{}).Error
}

func (e *employeeRepo) UpdateEmployee(ctx context.Context, emp model.Employee) error {
	return e.svcCtx.DbEngin.WithContext(ctx).Model(&emp).Where("id = ?", emp.ID).Updates(&emp).Error
}

func (e *employeeRepo) GetEmployee(ctx context.Context, Id int64) (model.Employee, error) {
	var emp model.Employee
	return emp, e.svcCtx.DbEngin.WithContext(ctx).Where("id = ?", Id).First(&emp).Error
}

func (e *employeeRepo) ListEmployee(ctx context.Context, req types.EmployeeListRequest) (int64, []model.Employee, error) {
	var emps []model.Employee
	var count int64
	order := "id"
	sort := "asc"
	db := e.svcCtx.DbEngin.WithContext(ctx)
	if req.Search != "" {
		db = db.Where("name like ? or company like ? or phone like ?", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%")
	}
	if req.Order != "" {
		order = req.Order
	}
	if req.Sort != "" {
		sort = req.Sort
	}

	return count, emps, db.Offset((int(req.Page) - 1) * int(req.PageSize)).
		Limit(int(req.PageSize)).
		Order(order + " " + sort).
		Find(&emps).Count(&count).Error
}
