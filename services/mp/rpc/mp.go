package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"wczero/services/mp/rpc/internal/config"
	"wczero/services/mp/rpc/internal/server"
	"wczero/services/mp/rpc/internal/svc"
	"wczero/services/mp/rpc/mp"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/mp.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)
	defer logx.Close()

	ctx := svc.NewServiceContext(c)
	srv := server.NewMPServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		mp.RegisterMPServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
