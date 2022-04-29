package logic

import (
	"context"
	"fmt"
	"time"

	"wczero/services/mp/api/internal/svc"
	"wczero/services/mp/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MPTextMsgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMPTextMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) MPTextMsgLogic {
	return MPTextMsgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MPTextMsgLogic) MPTextMsg(req types.MPTextMsgRequest) (*types.MPTextMsgResponse, error) {
	resp := &types.MPTextMsgResponse{
		ToUserName:   req.ToUserName,
		FromUserName: req.FromUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      fmt.Sprintf("[消息回复] - %s", time.Now().Format("2006-01-02 15:04:05")),
	}
	return resp, nil
}
