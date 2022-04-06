package logic

import (
	"context"
	"time"
	"wczero/common/jwtx"
	"wczero/services/user/rpc/user"

	"wczero/services/user/api/internal/svc"
	"wczero/services/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WeChatLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWeChatLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) WeChatLoginLogic {
	return WeChatLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WeChatLoginLogic) WeChatLogin(req types.WeChatLoginRequest) (*types.WeChatLoginResponse, error) {
	resp, err := l.svcCtx.UserRpc.WeChatLogin(l.ctx, &user.WeChatLoginRequest{
		Appid: req.AppID,
		Code:  req.Code,
	})
	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()

	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessSecret := l.svcCtx.Config.Auth.AccessSecret
	accessToken, err := jwtx.GetWeChatToken(accessSecret, now, accessExpire, resp.Openid)
	if err != nil {
		return nil, err
	}

	return &types.WeChatLoginResponse{
		AccessToken:  accessToken,
		AccessExpire: accessExpire,
	}, nil
}
