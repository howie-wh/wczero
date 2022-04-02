package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"wczero/services/user/api/internal/config"
	"wczero/services/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
