package employee

import (
	"context"

	repo "example/app/company/cmd/api/company/internal/repository"
	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteEmployeeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteEmployeeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteEmployeeLogic {
	return &DeleteEmployeeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteEmployeeLogic) DeleteEmployee(req *types.EmployeeDeleteRequest) (types.EmployeeDeleteResponse, error) {
	var resp types.EmployeeDeleteResponse
	resp.Msg = "删除成功"
	err := repo.NewEmployeeRepo(l.ctx, l.svcCtx).DeleteEmployee(l.ctx, req.Id)
	if err != nil {
		resp.Msg = "删除失败"
		return resp, err
	}
	return resp, nil
}
