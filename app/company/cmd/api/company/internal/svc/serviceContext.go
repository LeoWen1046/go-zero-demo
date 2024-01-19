package svc

import (
	"example/app/company/cmd/api/company/internal/config"
	"example/app/company/cmd/rpc/office/officeservice"
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	DbEngin         *gorm.DB
	OfficeRpcClient officeservice.OfficeService
}

func NewServiceContext(c config.Config, db *gorm.DB) *ServiceContext {
	return &ServiceContext{
		Config:          c,
		DbEngin:         db,
		OfficeRpcClient: officeservice.NewOfficeService(zrpc.MustNewClient(c.OfficeRpcConf.Rpc)),
	}
}
