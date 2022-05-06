package logic

import (
	"context"
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
		return nil, err
	}
	tList, tTotal, err := l.svcCtx.TypeModel.FindList(in.Start, in.Limit)
	if err != nil {
		return nil, err
	}

	var resp wallpaper.CategoryResponse
	for _, cl := range cList {
		resp.Category = append(resp.Category, cl.Category)
	}
	for _, tl := range tList {
		resp.Type = append(resp.Type, tl.Tp)
	}
	resp.CategoryTotal = cTotal
	resp.TypeTotal = tTotal

	return &resp, nil
}
