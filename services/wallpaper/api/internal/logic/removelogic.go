package logic

import (
	"context"
	"wczero/services/wallpaper/rpc/wallpaper"

	"wczero/services/wallpaper/api/internal/svc"
	"wczero/services/wallpaper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) RemoveLogic {
	return RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req types.RemoveRequest) (*types.RemoveResponse, error) {
	_, err := l.svcCtx.WallPaper.Remove(l.ctx, &wallpaper.RemoveRequest{
		List: req.List,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
