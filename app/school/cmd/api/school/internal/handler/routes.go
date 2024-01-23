// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	class "example/app/school/cmd/api/school/internal/handler/class"
	role "example/app/school/cmd/api/school/internal/handler/role"
	"example/app/school/cmd/api/school/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/class/create",
				Handler: class.CreateClassHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/class/get",
				Handler: class.GetClassHandler(serverCtx),
			},
		},
		rest.WithPrefix("/school/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/role/create",
				Handler: role.CreateRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/role/get",
				Handler: role.GetRoleHandler(serverCtx),
			},
		},
		rest.WithPrefix("/school/v1"),
	)
}