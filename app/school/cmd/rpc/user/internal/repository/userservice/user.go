package repository

import (
	"context"
	"example/app/school/cmd/rpc/user/internal/svc"
	"example/app/school/model"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user model.User) error
	DeleteUser(ctx context.Context, Id int64) error
	UpdateUser(ctx context.Context, user model.User) error
	GetUser(ctx context.Context, Id int64) (model.User, error)
	//ListOffice(ctx context.Context, req pb.) ([]*model.User, error)
}

type userRepo struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepo(c context.Context, serviceContext *svc.ServiceContext) *userRepo {
	return &userRepo{
		ctx:    c,
		svcCtx: serviceContext,
	}
}

func (o *userRepo) CreateUser(ctx context.Context, user model.User) error {
	return o.svcCtx.DbEngin.WithContext(ctx).Create(&user).Error
}

func (o *userRepo) DeleteUser(ctx context.Context, Id int64) error {
	return o.svcCtx.DbEngin.WithContext(ctx).Where("id = ?", Id).Delete(&model.User{}).Error
}

func (o *userRepo) UpdateUser(ctx context.Context, user model.User) error {
	return o.svcCtx.DbEngin.WithContext(ctx).Model(&user).Where("id = ?", user.ID).Updates(&user).Error
}

func (o *userRepo) GetUser(ctx context.Context, Id int64) (model.User, error) {
	var user model.User
	return user, o.svcCtx.DbEngin.WithContext(ctx).Where("id = ?", Id).First(&user).Error
}
