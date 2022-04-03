package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"wczero/services/wallpaper/api/internal/config"
	"wczero/services/wallpaper/api/internal/middleware"
	"wczero/services/wallpaper/rpc/wallpaperclient"
)

type ServiceContext struct {
	Config config.Config
	Error rest.Middleware
	WallPaper wallpaperclient.Wallpaper
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Error: middleware.NewErrorMiddleware().Handle,
		WallPaper: wallpaperclient.NewWallpaper(zrpc.MustNewClient(c.WallPaperRpc)),
	}
}
