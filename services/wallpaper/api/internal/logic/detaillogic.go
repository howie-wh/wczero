package logic

import (
	"context"
	"wczero/services/wallpaper/rpc/wallpaper"

	"wczero/services/wallpaper/api/internal/svc"
	"wczero/services/wallpaper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) DetailLogic {
	return DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req types.DetailRequest) (*types.DetailResponse, error) {
	resp, err := l.svcCtx.WallPaper.Detail(l.ctx, &wallpaper.DetailRequest{
		Wid: req.Wid,
		Start: req.Start,
		Limit: req.Limit,
	})
	if err != nil {
		return nil, err
	}

	apiResp := &types.DetailResponse{
		Total: resp.Total,
	}
	for _, v := range resp.List {
		wp := types.WallPaperInfo{
			Wid: v.Wid,
			Name: v.Name,
			ImageURL: v.ImageURL,
			Author: v.Author,
			Desc: v.Desc,
		}
		apiResp.List = append(apiResp.List, wp)
	}
	return apiResp, nil
}
