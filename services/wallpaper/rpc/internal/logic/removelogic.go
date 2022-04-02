package logic

import (
	"context"

	"wczero/services/wallpaper/rpc/internal/svc"
	"wczero/services/wallpaper/rpc/wallpaper"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *wallpaper.RemoveRequest) (*wallpaper.RemoveResponse, error) {
	for _, wid := range in.List {
		err := l.svcCtx.Model.DeleteByWid(wid)
		if err != nil {
			return nil, err
		}
	}
	return &wallpaper.RemoveResponse{}, nil
}
