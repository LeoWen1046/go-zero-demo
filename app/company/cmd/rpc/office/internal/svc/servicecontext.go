package svc

import (
	"example/app/company/cmd/rpc/office/internal/config"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DbEngin *gorm.DB
}

func NewServiceContext(c config.Config, db *gorm.DB) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DbEngin: db,
	}
}
