package logic

import (
	"context"
	"wczero/services/wallpaper/rpc/wallpaper"

	"wczero/services/wallpaper/api/internal/svc"
	"wczero/services/wallpaper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImportLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportLogic(ctx context.Context, svcCtx *svc.ServiceContext) ImportLogic {
	return ImportLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImportLogic) Import(req types.ImportRequest) (*types.ImportResponse, error) {
	rpcReq := &wallpaper.ImportRequest{}
	for _, v := range req.List {
		wp := &wallpaper.WallPaperInfo{
			Wid:      v.Wid,
			Name:     v.Name,
			Tid:      v.Tid,
			Cid:      v.Cid,
			ImageURL: v.ImageURL,
			Author:   v.Author,
			Desc:     v.Desc,
		}
		rpcReq.List = append(rpcReq.List, wp)
	}
	_, err := l.svcCtx.WallPaper.Import(l.ctx, rpcReq)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
