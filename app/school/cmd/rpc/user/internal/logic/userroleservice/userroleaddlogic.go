package userroleservicelogic

import (
	"context"
	repo "example/app/school/cmd/rpc/user/internal/repository/userroleservice"
	"example/app/school/model"

	"example/app/school/cmd/rpc/user/github.com/example/user"
	"example/app/school/cmd/rpc/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRoleAddLogic {
	return &UserRoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserRoleAddLogic) UserRoleAdd(in *user.UserRoleAddReq) (*user.UserRoleAddResp, error) {
	userRole := model.UserRole{
		Name: in.Name,
	}
	resp := new(user.UserRoleAddResp)
	resp.Msg = "创建成功"

	err := repo.NewUserRoleRepo(l.ctx, l.svcCtx).CreateUserRole(l.ctx, userRole)
	if err != nil {
		resp.Msg = "创建失败"
		return resp, err
	}

	return resp, nil
}
