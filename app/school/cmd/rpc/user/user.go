package main

import (
	"flag"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"example/app/school/cmd/rpc/user/github.com/example/user"
	"example/app/school/cmd/rpc/user/internal/config"
	userclassserviceServer "example/app/school/cmd/rpc/user/internal/server/userclassservice"
	userroleserviceServer "example/app/school/cmd/rpc/user/internal/server/userroleservice"
	userserviceServer "example/app/school/cmd/rpc/user/internal/server/userservice"
	"example/app/school/cmd/rpc/user/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 数据库连接
	DataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", c.MySQL.User,
		c.MySQL.Password, c.MySQL.Host, c.MySQL.Port, c.MySQL.Database, c.MySQL.Charset)
	MysqlDB, err := gorm.Open(mysql.Open(DataSource), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	ctx := svc.NewServiceContext(c, MysqlDB)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServiceServer(grpcServer, userserviceServer.NewUserServiceServer(ctx))
		user.RegisterUserRoleServiceServer(grpcServer, userroleserviceServer.NewUserRoleServiceServer(ctx))
		user.RegisterUserClassServiceServer(grpcServer, userclassserviceServer.NewUserClassServiceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
