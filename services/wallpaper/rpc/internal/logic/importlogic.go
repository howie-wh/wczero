package logic

import (
	"context"
	"wczero/common/utils"
	"wczero/services/wallpaper/model"

	"wczero/services/wallpaper/rpc/internal/svc"
	"wczero/services/wallpaper/rpc/wallpaper"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	_minWidLen = 4
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

func (l *ImportLogic) WidGenerate(in *wallpaper.ImportRequest) error {
	tNum, err := l.svcCtx.NoCacheModel.GetTableCount()
	if err != nil {
		logx.Errorf("GetTableCount err: %v", err)
		return err
	}

	for _, wp := range in.List {
		widStr := utils.NumToBHex(tNum)
		for i := _minWidLen - len(widStr); i > 0; i-- { // set default char "0" if string len less 4
			widStr = "0" + widStr
		}

		wp.Wid = widStr
		tNum = tNum + 1
	}
	return nil
}

func (l *ImportLogic) Import(in *wallpaper.ImportRequest) (*wallpaper.ImportResponse, error) {
	if err := l.WidGenerate(in); err != nil {
		return nil, err
	}

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
