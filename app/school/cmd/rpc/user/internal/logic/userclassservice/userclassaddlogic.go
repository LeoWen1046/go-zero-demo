package userclassservicelogic

import (
	"context"
	"example/app/school/cmd/rpc/user/github.com/example/user"
	repo "example/app/school/cmd/rpc/user/internal/repository/userclassservice"
	"example/app/school/cmd/rpc/user/internal/svc"
	"example/app/school/model"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserClassAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserClassAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserClassAddLogic {
	return &UserClassAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserClassAddLogic) UserClassAdd(in *user.UserClassAddReq) (*user.UserClassAddResp, error) {
	userClass := model.UserClass{
		Name: in.Name,
		Room: in.Room,
	}
	resp := new(user.UserClassAddResp)
	resp.Msg = "创建成功"

	err := repo.NewUserClassRepo(l.ctx, l.svcCtx).CreateUserClass(l.ctx, userClass)
	if err != nil {
		resp.Msg = "创建失败"
		return resp, err
	}

	return resp, nil
}
