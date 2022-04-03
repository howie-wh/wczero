package handler

import (
	"net/http"
	"wczero/common/response" //1

	"wczero/services/user/api/internal/logic"
	"wczero/services/user/api/internal/svc"
)

func AdminUserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAdminUserInfoLogic(r.Context(), svcCtx)
		resp, err := l.AdminUserInfo()
		response.Response(w, resp, err) //2
	}
}
