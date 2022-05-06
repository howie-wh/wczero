package logic

import (
	"context"
	"wczero/services/wallpaper/rpc/wallpaper"

	"wczero/services/wallpaper/api/internal/svc"
	"wczero/services/wallpaper/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) CategoryLogic {
	return CategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryLogic) Category(req types.CategoryRequest) (*types.CategoryResponse, error) {
	logx.Infof("CategoryRequest : %v\n", req)
	resp, err := l.svcCtx.WallPaper.Category(l.ctx, &wallpaper.CategoryRequest{
		Start: req.Start,
		Limit: req.Limit,
	})
	if err != nil {
		logx.Errorf("get Category, err:%v\n", err)
		return nil, err
	}

	return &types.CategoryResponse{
		Category:      resp.Category,
		CategoryTotal: resp.CategoryTotal,
		Tp:            resp.Tp,
		TpTotal:       resp.TpTotal,
	}, nil
}
