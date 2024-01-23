package userclassservicelogic

import (
	"context"
	"example/app/school/cmd/rpc/user/github.com/example/user"
	repo "example/app/school/cmd/rpc/user/internal/repository/userclassservice"
	"example/app/school/cmd/rpc/user/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserClassInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserClassInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserClassInfoLogic {
	return &UserClassInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserClassInfoLogic) UserClassInfo(in *user.UserClassInfoReq) (*user.UserClassInfoResp, error) {
	resp := new(user.UserClassInfoResp)
	userClass, getErr := repo.NewUserClassRepo(l.ctx, l.svcCtx).GetUserClass(l.ctx, in.Id)
	if getErr != nil {
		return resp, getErr
	}

	resp.UserClass = &user.UserClass{
		Id:   in.Id,
		Name: userClass.Name,
		Room: userClass.Room,
	}

	return resp, nil
}
