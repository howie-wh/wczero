package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/logx"

	"wczero/services/wallpaper/rpc/internal/config"
	"wczero/services/wallpaper/rpc/internal/server"
	"wczero/services/wallpaper/rpc/internal/svc"
	"wczero/services/wallpaper/rpc/wallpaper"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/wallpaper.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)
	defer logx.Close()

	ctx := svc.NewServiceContext(c)
	srv := server.NewWallpaperServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		wallpaper.RegisterWallpaperServer(grpcServer, srv)

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	logx.Infof("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
