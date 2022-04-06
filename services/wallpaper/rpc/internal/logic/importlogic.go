package logic

import (
	"context"
	"wczero/services/wallpaper/model"

	"wczero/services/wallpaper/rpc/internal/svc"
	"wczero/services/wallpaper/rpc/wallpaper"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewImportLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportLogic {
	return &ImportLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ImportLogic) Import(in *wallpaper.ImportRequest) (*wallpaper.ImportResponse, error) {
	tabList := make([]*model.WallpaperTab, 0)
	for _, wp := range in.List {
		tab := &model.WallpaperTab{
			Wid:      wp.Wid,
			Name:     wp.Name,
			Category: wp.Category,
			ImageUrl: wp.ImageURL,
			Author:   wp.Author,
			Desc:     wp.Desc,
		}
		tabList = append(tabList, tab)
	}
	err := l.svcCtx.NoCacheModel.BulkInsert(tabList)
	if err != nil {
		return nil, err
	}
	return &wallpaper.ImportResponse{}, nil
}
