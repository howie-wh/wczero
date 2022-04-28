package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"wczero/services/mp/api/internal/config"
	"wczero/services/mp/rpc/mpclient"
)

type ServiceContext struct {
	Config config.Config
	MP     mpclient.MP
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		MP:     mpclient.NewMP(zrpc.MustNewClient(c.MPRpc)),
	}
}
