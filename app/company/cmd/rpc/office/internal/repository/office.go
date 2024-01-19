package repository

import (
	"context"
	"example/app/company/cmd/rpc/office/internal/svc"
	"example/app/company/cmd/rpc/office/pb"
	"example/app/company/model"
)

type OfficeRepo interface {
	CreateOffice(ctx context.Context, office model.Office) error
	DeleteOffice(ctx context.Context, Id int64) error
	UpdateOffice(ctx context.Context, office model.Office) error
	GetOffice(ctx context.Context, Id int64) (model.Office, error)
	ListOffice(ctx context.Context, req pb.ListOfficeRequest) ([]*model.Office, error)
}

type officeRepo struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOfficeRepo(c context.Context, serviceContext *svc.ServiceContext) *officeRepo {
	return &officeRepo{
		ctx:    c,
		svcCtx: serviceContext,
	}
}

func (o *officeRepo) CreateOffice(ctx context.Context, office model.Office) error {
	return o.svcCtx.DbEngin.WithContext(ctx).Create(&office).Error
}

func (o *officeRepo) DeleteOffice(ctx context.Context, Id int64) error {
	return o.svcCtx.DbEngin.WithContext(ctx).Where("id = ?", Id).Delete(&model.Office{}).Error
}

func (o *officeRepo) UpdateOffice(ctx context.Context, office model.Office) error {
	return o.svcCtx.DbEngin.WithContext(ctx).Model(&office).Where("id = ?", office.ID).Updates(&office).Error
}

func (o *officeRepo) GetOffice(ctx context.Context, Id int64) (model.Office, error) {
	var office model.Office
	return office, o.svcCtx.DbEngin.WithContext(ctx).Where("id = ?", Id).First(&office).Error
}

func (o *officeRepo) ListOffice(ctx context.Context, req pb.ListOfficeRequest) (int64, []model.Office, error) {
	var offices []model.Office
	var count int64
	order := "id"
	sort := "asc"
	db := o.svcCtx.DbEngin.WithContext(ctx)
	if req.Search != "" {
		db = db.Where("name like ? or address like ?", "%"+req.Search+"%", "%"+req.Search+"%")
	}
	if req.Order != "" {
		order = req.Order
	}
	if req.Sort != "" {
		sort = req.Sort
	}

	return count, offices, db.Offset((int(req.Page) - 1) * int(req.PageSize)).
		Limit(int(req.PageSize)).
		Order(order + " " + sort).
		Find(&offices).Count(&count).Error
}
