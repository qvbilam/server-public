package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"public/api"
	proto "public/api/qvbilam/public/v1"
	"public/global"
	"public/initialize"
)

func main() {
	// 初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDatabase()
	initialize.InitServer()

	// 注册服务
	server := grpc.NewServer()
	// todo
	proto.RegisterVideoServer(server, &api.VideoServer{})
	proto.RegisterFileServer(server, &api.FileServer{})

	Host := "0.0.0.0"
	Port := 9803

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", Host, Port))
	if err != nil {
		zap.S().Panicf("listen port error: %s", err)
	}

	zap.S().Infof("start %s server, host: %s:%d", global.ServerConfig.Name, Host, Port)
	go func() {
		if err := server.Serve(lis); err != nil {
			zap.S().Panicf("start server error: %s", err)
		}
	}()

	// 监听结束
	quit := make(chan os.Signal)
	<-quit
}
