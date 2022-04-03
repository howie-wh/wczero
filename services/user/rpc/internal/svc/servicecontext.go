package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"wczero/services/user/model"
	"wczero/services/user/rpc/internal/config"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.UserTabModel
	AdminModel model.UserAdminTabModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		UserModel:  model.NewUserTabModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
		AdminModel: model.NewUserAdminTabModel(sqlx.NewMysql(c.Mysql.DataSource), c.CacheRedis),
	}
}
