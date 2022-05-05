package logic

import (
	"context"
	"wczero/services/mp/api/internal/svc"
	"wczero/services/mp/api/internal/types"
	"wczero/services/mp/rpc/mp"

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
	resp, err := l.svcCtx.MP.MPTextMsg(l.ctx, &mp.MPTextMsgRequest{
		ToUserName:   req.ToUserName,
		FromUserName: req.FromUserName,
		CreateTime:   req.CreateTime,
		MsgType:      req.MsgType,
		Content:      req.Content,
		MsgId:        req.MsgId,
	})
	if err != nil {
		return nil, err
	}

	apiResp := types.MPTextMsgResponse{
		ToUserName:   resp.ToUserName,
		FromUserName: resp.FromUserName,
		CreateTime:   resp.CreateTime,
		MsgType:      resp.MsgType,
		Content:      resp.Content,
	}
	return &apiResp, nil
}
