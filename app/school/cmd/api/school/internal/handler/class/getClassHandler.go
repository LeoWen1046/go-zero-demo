package class

import (
	"example/common/result"
	"net/http"

	"example/app/school/cmd/api/school/internal/logic/class"
	"example/app/school/cmd/api/school/internal/svc"
	"example/app/school/cmd/api/school/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetClassHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserClassDetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := class.NewGetClassLogic(r.Context(), svcCtx)
		resp, err := l.GetClass(&req)
		result.HttpResult(r, w, resp, err)
	}
}
