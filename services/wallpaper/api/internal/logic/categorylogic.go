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

	apiResp := types.CategoryResponse{
		TpTotal:       resp.TpTotal,
		CategoryTotal: resp.CategoryTotal,
	}
	for _, tp := range resp.Tp {
		tpInfo := types.TypeInfo{
			Tid:     tp.Tid,
			Name:    tp.Name,
			CidList: tp.CidList,
		}
		apiResp.Tp = append(apiResp.Tp, tpInfo)
	}
	for _, category := range resp.Category {
		categoryInfo := types.CategoryInfo{
			Cid:  category.Cid,
			Name: category.Name,
		}
		apiResp.Category = append(apiResp.Category, categoryInfo)
	}

	return &apiResp, nil
}
