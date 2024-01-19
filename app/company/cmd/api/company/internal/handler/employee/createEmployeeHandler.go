package employee

import (
	"example/common/result"
	"net/http"

	"example/app/company/cmd/api/company/internal/logic/employee"
	"example/app/company/cmd/api/company/internal/svc"
	"example/app/company/cmd/api/company/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateEmployeeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EmployeeCreateRequest
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := employee.NewCreateEmployeeLogic(r.Context(), svcCtx)
		resp, err := l.CreateEmployee(&req)
		result.HttpResult(r, w, resp, err)
	}
}
