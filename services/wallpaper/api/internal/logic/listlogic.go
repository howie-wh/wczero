package logic

import (
	"context"
	"wczero/services/wallpaper/rpc/wallpaper"

	"wczero/services/wallpaper/api/internal/svc"
	"wczero/services/wallpaper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListLogic {
	return ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req types.ListRequest) (*types.ListResponse, error) {
	resp, err := l.svcCtx.WallPaper.List(l.ctx, &wallpaper.ListRequest{
		Start: req.Start,
		Limit: req.Limit,
	})
	if err != nil {
		return nil, err
	}

	apiResp := &types.ListResponse{
		Total: resp.Total,
	}
	for _, v := range resp.List {
		wp := types.WallPaperInfo{
			Wid:      v.Wid,
			Name:     v.Name,
			ImageURL: v.ImageURL,
			Author:   v.Author,
			Desc:     v.Desc,
		}
		apiResp.List = append(apiResp.List, wp)
	}
	return apiResp, nil
}
