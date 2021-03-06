package config

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
		Table      string
	}
	CacheRedis cache.CacheConf
	Salt       string
	AppSecret  string
	Log        logx.LogConf
	QiNiu      struct {
		Domain string
		Zone   string
	}
}
