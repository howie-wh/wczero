package handler

import (
	"net/http"
	"wczero/common/response" //1

	"github.com/zeromicro/go-zero/rest/httpx"
	"wczero/services/user/api/internal/logic"
	"wczero/services/user/api/internal/svc"
	"wczero/services/user/api/internal/types"
)

func WeChatLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WeChatLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewWeChatLoginLogic(r.Context(), svcCtx)
		resp, err := l.WeChatLogin(req)
		response.Response(w, resp, err) //2
	}
}
