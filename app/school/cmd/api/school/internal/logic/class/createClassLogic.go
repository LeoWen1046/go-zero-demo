package class

import (
	"context"
	"example/app/school/cmd/api/school/internal/svc"
	"example/app/school/cmd/api/school/internal/types"
	"example/app/school/cmd/rpc/user/github.com/example/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateClassLogic {
	return &CreateClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateClassLogic) CreateClass(req *types.UserClassCreateRequest) (types.UserClassCreateResponse, error) {
	var resp types.UserClassCreateResponse

	classResp, err := l.svcCtx.UserClassRpcClient.UserClassAdd(l.ctx, &user.UserClassAddReq{
		Name: req.Name,
		Room: req.Room,
	})

	resp.Msg = classResp.Msg

	return resp, err
}
