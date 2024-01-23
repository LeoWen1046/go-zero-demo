package svc

import (
	"example/app/school/cmd/api/school/internal/config"
	"example/app/school/cmd/rpc/user/client/userclassservice"
	"example/app/school/cmd/rpc/user/client/userroleservice"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config             config.Config
	DbEngin            *gorm.DB
	UserRoleRpcClient  userroleservice.UserRoleService
	UserClassRpcClient userclassservice.UserClassService
}

func NewServiceContext(c config.Config, db *gorm.DB) *ServiceContext {
	return &ServiceContext{
		Config:             c,
		DbEngin:            db,
		UserRoleRpcClient:  userroleservice.NewUserRoleService(zrpc.MustNewClient(c.SchoolRpcConf.Rpc)),
		UserClassRpcClient: userclassservice.NewUserClassService(zrpc.MustNewClient(c.SchoolRpcConf.Rpc)),
	}
}
