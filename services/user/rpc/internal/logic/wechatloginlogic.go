package logic

import (
	"context"

	"wczero/services/user/rpc/internal/svc"
	"wczero/services/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type WeChatLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewWeChatLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WeChatLoginLogic {
	return &WeChatLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *WeChatLoginLogic) WeChatLogin(in *user.WeChatLoginRequest) (*user.WeChatLoginResponse, error) {

	return &user.WeChatLoginResponse{Token: "wechat login successful..."}, nil
}
