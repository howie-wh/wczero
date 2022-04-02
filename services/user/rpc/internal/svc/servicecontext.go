package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"wczero/services/user/model"
	"wczero/services/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Model model.UserTabModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model: model.NewUserTabModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
