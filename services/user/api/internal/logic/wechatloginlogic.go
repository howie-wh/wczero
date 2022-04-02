package logic

import (
	"context"
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
	_, err := l.svcCtx.User.WeChatLogin(l.ctx, &user.WeChatLoginRequest{
		Appid: req.AppID,
		Code: req.Code,
	})
	if err != nil {
		return nil, err
	}

	return &types.WeChatLoginResponse{}, nil
}
