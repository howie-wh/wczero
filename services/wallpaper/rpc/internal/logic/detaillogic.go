package logic

import (
	"context"
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
	resp, err := l.svcCtx.Model.FindOneByWid(in.Wid)
	if err != nil {
		return nil, err
	}

	detailResp := &wallpaper.DetailResponse{
		Wid:      resp.Wid,
		Name:     resp.Name,
		Type:     resp.Tp,
		Category: resp.Category,
		ImageURL: resp.ImageUrl,
		Author:   resp.Author,
		Desc:     resp.Desc,
	}
	return detailResp, nil
}
