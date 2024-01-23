package class

import (
	"context"
	"example/app/school/cmd/rpc/user/github.com/example/user"

	"example/app/school/cmd/api/school/internal/svc"
	"example/app/school/cmd/api/school/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetClassLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetClassLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetClassLogic {
	return &GetClassLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetClassLogic) GetClass(req *types.UserClassDetailRequest) (types.UserClassDetailResponse, error) {
	var resp types.UserClassDetailResponse

	classResp, err := l.svcCtx.UserClassRpcClient.UserClassInfo(l.ctx, &user.UserClassInfoReq{
		Id: req.ID,
	})

	resp.UserClass = types.UserClass{
		ID:   classResp.UserClass.Id,
		Name: classResp.UserClass.Name,
		Room: classResp.UserClass.Room,
	}

	return resp, err
}
