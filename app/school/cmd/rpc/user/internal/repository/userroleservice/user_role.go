package repository

import (
	"context"
	"example/app/school/cmd/rpc/user/internal/svc"
	"example/app/school/model"
)

type UserRoleRepo interface {
	CreateUserRole(ctx context.Context, userRole model.UserRole) error
	DeleteUserRole(ctx context.Context, Id int64) error
	UpdateUserRole(ctx context.Context, userRole model.UserRole) error
	GetUserRole(ctx context.Context, Id int64) (model.UserRole, error)
	//ListOffice(ctx context.Context, req pb.ListOfficeRequest) ([]*model.Office, error)
}

type userRoleRepo struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRoleRepo(c context.Context, serviceContext *svc.ServiceContext) *userRoleRepo {
	return &userRoleRepo{
		ctx:    c,
		svcCtx: serviceContext,
	}
}

func (o *userRoleRepo) CreateUserRole(ctx context.Context, userRole model.UserRole) error {
	return o.svcCtx.DbEngin.WithContext(ctx).Create(&userRole).Error
}

func (o *userRoleRepo) DeleteUserRole(ctx context.Context, Id int64) error {
	return o.svcCtx.DbEngin.WithContext(ctx).Where("id = ?", Id).Delete(&model.UserRole{}).Error
}

func (o *userRoleRepo) UpdateUserRole(ctx context.Context, userRole model.UserRole) error {
	return o.svcCtx.DbEngin.WithContext(ctx).Model(&userRole).Where("id = ?", userRole.ID).Updates(&userRole).Error
}

func (o *userRoleRepo) GetUserRole(ctx context.Context, Id int64) (model.UserRole, error) {
	var userRole model.UserRole
	return userRole, o.svcCtx.DbEngin.WithContext(ctx).Where("id = ?", Id).First(&userRole).Error
}
