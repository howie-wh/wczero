package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"wczero/services/user/api/internal/config"
	"wczero/services/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	User userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
