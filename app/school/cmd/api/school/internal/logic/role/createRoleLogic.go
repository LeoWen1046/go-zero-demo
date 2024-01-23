package role

import (
	"context"
	"example/app/school/cmd/rpc/user/github.com/example/user"

	"example/app/school/cmd/api/school/internal/svc"
	"example/app/school/cmd/api/school/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRoleLogic {
	return &CreateRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRoleLogic) CreateRole(req *types.UserRoleCreateRequest) (types.UserRoleCreateResponse, error) {
	var resp types.UserRoleCreateResponse

	roleResp, err := l.svcCtx.UserRoleRpcClient.UserRoleAdd(l.ctx, &user.UserRoleAddReq{
		Name: req.Name,
	})

	resp.Msg = roleResp.Msg

	return resp, err
}
