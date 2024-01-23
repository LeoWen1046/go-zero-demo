package class

import (
	"example/common/result"
	"net/http"

	"example/app/school/cmd/api/school/internal/logic/class"
	"example/app/school/cmd/api/school/internal/svc"
	"example/app/school/cmd/api/school/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateClassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserClassCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := class.NewCreateClassLogic(r.Context(), svcCtx)
		resp, err := l.CreateClass(&req)
		result.HttpResult(r, w, resp, err)
	}
}
