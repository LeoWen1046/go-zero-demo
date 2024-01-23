package userroleservicelogic

import (
	"context"
	repo "example/app/school/cmd/rpc/user/internal/repository/userroleservice"

	"example/app/school/cmd/rpc/user/github.com/example/user"
	"example/app/school/cmd/rpc/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRoleInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleInfoLogic {
	return &UserRoleInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRoleInfoLogic) UserRoleInfo(in *user.UserRoleInfoReq) (*user.UserRoleInfoResp, error) {
	resp := new(user.UserRoleInfoResp)
	userRole, getErr := repo.NewUserRoleRepo(l.ctx, l.svcCtx).GetUserRole(l.ctx, in.Id)
	if getErr != nil {
		return resp, getErr
	}

	resp.UserRole = &user.UserRole{
		Id:   in.Id,
		Name: userRole.Name,
	}

	return resp, nil
}
