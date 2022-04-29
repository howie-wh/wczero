package handler

import (
	"encoding/xml"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"wczero/services/mp/api/internal/logic"
	"wczero/services/mp/api/internal/svc"
	"wczero/services/mp/api/internal/types"
)

type MPReplyTextMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int64  `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	// 若不标记XMLName, 则解析后的xml名为该结构体的名称
	XMLName xml.Name `xml:"xml"`
}

func MPTextMsgHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MPTextMsgRequest
		if err := httpx.Parse(r, &req); err != nil {
			logx.Errorf("http.Parse err: %v\n", err)
			httpx.Error(w, err)
			return
		}

		l := logic.NewMPTextMsgLogic(r.Context(), svcCtx)
		resp, err := l.MPTextMsg(req)
		if err != nil {
			logx.Errorf("MPTextMsg err: %v\n", err)
			httpx.Error(w, err)
			return
		}

		msgResp := &MPReplyTextMsg{
			ToUserName:   resp.ToUserName,
			FromUserName: resp.FromUserName,
			CreateTime:   resp.CreateTime,
			MsgType:      resp.MsgType,
			Content:      resp.Content,
		}
		msg, err := xml.Marshal(msgResp)
		if err != nil {
			logx.Errorf("xml marshal err: %v\n", err)
			httpx.Error(w, err)
			return
		}
		httpx.OkJson(w, msg)
	}
}
