package role

import (
	"context"
	"example/app/school/cmd/rpc/user/github.com/example/user"

	"example/app/school/cmd/api/school/internal/svc"
	"example/app/school/cmd/api/school/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRoleLogic {
	return &GetRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleLogic) GetRole(req *types.UserRoleDetailRequest) (types.UserRoleDetailResponse, error) {
	var resp types.UserRoleDetailResponse

	roleResp, err := l.svcCtx.UserRoleRpcClient.UserRoleInfo(l.ctx, &user.UserRoleInfoReq{
		Id: req.ID,
	})

	resp.UserRole = types.UserRole{
		ID:   roleResp.UserRole.Id,
		Name: roleResp.UserRole.Name,
	}
	return resp, err
}
