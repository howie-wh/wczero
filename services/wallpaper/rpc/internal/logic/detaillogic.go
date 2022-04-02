package logic

import (
	"context"
	"wczero/services/wallpaper/model"

	"wczero/services/wallpaper/rpc/internal/svc"
	"wczero/services/wallpaper/rpc/wallpaper"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *wallpaper.DetailRequest) (*wallpaper.DetailResponse, error) {
	wallpaperList := make([]*model.WallpaperTab, 0)
	if in.Wid != "" {
		resp, err := l.svcCtx.Model.FindOneByWid(in.Wid)
		if err != nil {
			return nil,err
		}
		wallpaperList = append(wallpaperList, resp)
	} else {
		resp, err := l.svcCtx.Model.FindList(in.Start, in.Limit)
		if err != nil {
			return nil,err
		}
		wallpaperList = append(wallpaperList, resp...)
	}

	var detailResp wallpaper.DetailResponse
	for _, wpt := range wallpaperList {
		wp := &wallpaper.WallPaperInfo{
			Wid: wpt.Wid,
			Name: wpt.Name,
			ImageURL: wpt.ImageUrl,
			Author: wpt.Author,
			Desc: wpt.Desc,
		}
		detailResp.List = append(detailResp.List, wp)
	}
	return &detailResp, nil
}
