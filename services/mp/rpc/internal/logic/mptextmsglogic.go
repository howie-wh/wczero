package logic

import (
	"context"
	"fmt"
	"time"
	"wczero/services/mp/rpc/internal/svc"
	"wczero/services/mp/rpc/mp"

	"github.com/zeromicro/go-zero/core/logx"
)

var (
	_defaultTextMsgReply = fmt.Sprintf("[消息回复] - %s", time.Now().Format("2006-01-02 15:04:05"))
	_textMsgType         = "text"
	_textContextLen      = 4
)

type MPTextMsgLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMPTextMsgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MPTextMsgLogic {
	return &MPTextMsgLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MPTextMsgLogic) ValidationTextMsg(in *mp.MPTextMsgRequest) bool {
	if in.MsgType != _textMsgType {
		return false
	}
	if len(in.Content) < _textContextLen {
		return false
	}
	return true
}

func (l *MPTextMsgLogic) MPTextMsg(in *mp.MPTextMsgRequest) (*mp.MPTextMsgResponse, error) {
	resp := mp.MPTextMsgResponse{
		ToUserName:   in.FromUserName,
		FromUserName: in.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      _textMsgType,
		Content:      _defaultTextMsgReply,
	}
	if !l.ValidationTextMsg(in) {
		logx.Errorf("param invalid...")
		return &resp, nil
	}

	respWallpaper, err := l.svcCtx.Model.FindOneByWid(in.Content)
	if err != nil {
		logx.Errorf("sql query failed, err: %v", err)
		return &resp, nil
	}

	logx.Info(respWallpaper.ImageUrl)
	resp.Content = respWallpaper.ImageUrl
	return &resp, nil
}
