package logic

import (
	"context"
	"wczero/services/user/rpc/userclient"

	"wczero/services/user/api/internal/svc"
	"wczero/services/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) AdminRegisterLogic {
	return AdminRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminRegisterLogic) AdminRegister(req types.AdminRegisterRequest) (resp *types.AdminRegisterResponse, err error) {
	res, err := l.svcCtx.UserRpc.AdminRegister(l.ctx, &userclient.AdminRegisterRequest{
		Name:     req.Name,
		Gender:   req.Gender,
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &types.AdminRegisterResponse{
		Id:     res.Id,
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
