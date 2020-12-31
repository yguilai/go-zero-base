// Code generated by goctl. DO NOT EDIT!
// Source: captcha.proto

package main

import (
	"flag"
	"fmt"

	"github.com/yguilai/timetable-micro/services/captcha/rpc/captcha"
	"github.com/yguilai/timetable-micro/services/captcha/rpc/internal/config"
	"github.com/yguilai/timetable-micro/services/captcha/rpc/internal/server"
	"github.com/yguilai/timetable-micro/services/captcha/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/captcha.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewCaptchaServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		captcha.RegisterCaptchaServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}