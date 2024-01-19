// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	employee "example/app/company/cmd/api/company/internal/handler/employee"
	office "example/app/company/cmd/api/company/internal/handler/office"
	"example/app/company/cmd/api/company/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/employee/create",
				Handler: employee.CreateEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/employee/detele",
				Handler: employee.DeleteEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/employee/update/:id",
				Handler: employee.UpdateEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/employee/get",
				Handler: employee.GetEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/employee/list",
				Handler: employee.ListEmployeeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/company/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/office/create",
				Handler: office.CreateOfficeHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/office/detele",
				Handler: office.DeleteOfficeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/office/update/:id",
				Handler: office.UpdateOfficeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/office/get",
				Handler: office.GetOfficeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/office/list",
				Handler: office.ListOfficeHandler(serverCtx),
			},
		},
		rest.WithPrefix("/company/v1"),
	)
}
