package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"wczero/services/wallpaper/model"
	"wczero/services/wallpaper/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	Model         model.WallpaperTabModel // 手动代码
	NoCacheModel  model.NoCacheWallpaperTabModel
	CategoryModel model.WallpaperCategoryTabModel
	TypeModel     model.WallpaperTypeTabModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		Model:         model.NewWallpaperTabModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis), // 手动代码
		NoCacheModel:  model.NewNoCacheWallpaperTabModel(sqlx.NewMysql(c.Mysql.DataSource)),
		CategoryModel: model.NewWallpaperCategoryTabModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
		TypeModel:     model.NewWallpaperTypeTabModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
