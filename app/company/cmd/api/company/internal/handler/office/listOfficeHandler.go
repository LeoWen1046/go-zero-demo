package office

import (
	"example/common/result"
	"net/http"

	"example/app/company/cmd/api/company/internal/logic/office"
	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListOfficeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OfficeListRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := office.NewListOfficeLogic(r.Context(), svcCtx)
		resp, err := l.ListOffice(req)
		result.HttpResult(r, w, resp, err)
	}
}
