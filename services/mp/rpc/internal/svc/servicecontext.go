package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"wczero/services/mp/rpc/internal/config"
	"wczero/services/wallpaper/model"
)

type ServiceContext struct {
	Config config.Config
	Model  model.WallpaperTabModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewWallpaperTabModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
