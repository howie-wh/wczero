package logic

import (
	"context"
	"wczero/services/wallpaper/rpc/internal/svc"
	"wczero/services/wallpaper/rpc/wallpaper"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *wallpaper.ListRequest) (*wallpaper.ListResponse, error) {
	resp, total, err := l.svcCtx.Model.FindList(in.Start, in.Limit)
	if err != nil {
		return nil, err
	}

	var detailResp wallpaper.ListResponse
	for _, wpt := range resp {
		wp := &wallpaper.WallPaperInfo{
			Wid:      wpt.Wid,
			Name:     wpt.Name,
			ImageURL: wpt.ImageUrl,
			Author:   wpt.Author,
			Desc:     wpt.Desc,
		}
		detailResp.List = append(detailResp.List, wp)
	}
	detailResp.Total = total
	return &detailResp, nil
}
