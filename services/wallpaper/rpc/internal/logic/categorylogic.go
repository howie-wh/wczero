package logic

import (
	"context"
	"encoding/json"
	"wczero/services/wallpaper/rpc/internal/svc"
	"wczero/services/wallpaper/rpc/wallpaper"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryLogic {
	return &CategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CategoryLogic) Category(in *wallpaper.CategoryRequest) (*wallpaper.CategoryResponse, error) {
	cList, cTotal, err := l.svcCtx.CategoryModel.FindList(in.Start, in.Limit)
	if err != nil {
		logx.Errorf("find category list error: %v\n", err)
		return nil, err
	}
	tList, tTotal, err := l.svcCtx.TypeModel.FindList(in.Start, in.Limit)
	if err != nil {
		logx.Errorf("find type list error: %v\n", err)
		return nil, err
	}

	var resp wallpaper.CategoryResponse
	for _, cl := range cList {
		categoryInfo := &wallpaper.CategoryInfo{
			Cid:  cl.Id,
			Name: cl.Category,
		}
		resp.Category = append(resp.Category, categoryInfo)
	}
	for _, tl := range tList {
		tpInfo := &wallpaper.TypeInfo{
			Tid:  tl.Id,
			Name: tl.Tp,
		}
		_ = json.Unmarshal([]byte(tl.CidList), &tpInfo.CidList)
		resp.Tp = append(resp.Tp, tpInfo)
	}
	resp.CategoryTotal = cTotal
	resp.TpTotal = tTotal

	logx.Infof("Category Response: %v\n", resp)
	return &resp, nil
}
