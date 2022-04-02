package handler

import (
	"net/http"
	"wczero/common/response" //1

	"github.com/zeromicro/go-zero/rest/httpx"
	"wczero/services/user/api/internal/logic"
	"wczero/services/user/api/internal/svc"
	"wczero/services/user/api/internal/types"
)

func AdminLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AdminLoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewAdminLoginLogic(r.Context(), svcCtx)
		resp, err := l.AdminLogin(req)
		response.Response(w, resp, err) //2
	}
}
