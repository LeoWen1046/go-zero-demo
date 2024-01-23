package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	MySQL struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
		Charset  string
	}
	SchoolRpcConf struct {
		Rpc zrpc.RpcClientConf
	}
}
