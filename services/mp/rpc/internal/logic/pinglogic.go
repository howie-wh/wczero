package logic

import (
	"context"

	"wczero/services/mp/rpc/internal/svc"
	"wczero/services/mp/rpc/mp"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *mp.Request) (*mp.Response, error) {
	// todo: add your logic here and delete this line

	return &mp.Response{}, nil
}
