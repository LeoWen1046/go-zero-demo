package role

import (
	"example/common/result"
	"net/http"

	"example/app/school/cmd/api/school/internal/logic/role"
	"example/app/school/cmd/api/school/internal/svc"
	"example/app/school/cmd/api/school/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRoleDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewGetRoleLogic(r.Context(), svcCtx)
		resp, err := l.GetRole(&req)
		result.HttpResult(r, w, resp, err)
	}
}
