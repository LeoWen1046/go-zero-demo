package repository

import (
	"context"
	"example/app/school/cmd/rpc/user/internal/svc"
	"example/app/school/model"
)

type UserClassRepo interface {
	CreateUserClass(ctx context.Context, office model.UserClass) error
	DeleteUserClass(ctx context.Context, Id int64) error
	UpdateUserClass(ctx context.Context, office model.UserClass) error
	GetUserClass(ctx context.Context, Id int64) (model.UserClass, error)
	//ListOffice(ctx context.Context, req pb.ListOfficeRequest) ([]*model.Office, error)
}

type userClassRepo struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserClassRepo(c context.Context, serviceContext *svc.ServiceContext) *userClassRepo {
	return &userClassRepo{
		ctx:    c,
		svcCtx: serviceContext,
	}
}

func (o *userClassRepo) CreateUserClass(ctx context.Context, userClass model.UserClass) error {
	return o.svcCtx.DbEngin.WithContext(ctx).Create(&userClass).Error
}

func (o *userClassRepo) DeleteUserClass(ctx context.Context, Id int64) error {
	return o.svcCtx.DbEngin.WithContext(ctx).Where("id = ?", Id).Delete(&model.UserClass{}).Error
}

func (o *userClassRepo) UpdateUserClass(ctx context.Context, userClass model.UserClass) error {
	return o.svcCtx.DbEngin.WithContext(ctx).Model(&userClass).Where("id = ?", userClass.ID).Updates(&userClass).Error
}

func (o *userClassRepo) GetUserClass(ctx context.Context, Id int64) (model.UserClass, error) {
	var userClass model.UserClass
	return userClass, o.svcCtx.DbEngin.WithContext(ctx).Where("id = ?", Id).First(&userClass).Error
}
