package logic

import (
	"context"
	"time"
	"wczero/common/jwtx"
	"wczero/services/user/rpc/userclient"

	"wczero/services/user/api/internal/svc"
	"wczero/services/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) AdminLoginLogic {
	return AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req types.AdminLoginRequest) (resp *types.AdminLoginResponse, err error) {
	res, err := l.svcCtx.UserRpc.AdminLogin(l.ctx, &userclient.AdminLoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}

	return &types.AdminLoginResponse{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
	}, nil
}
