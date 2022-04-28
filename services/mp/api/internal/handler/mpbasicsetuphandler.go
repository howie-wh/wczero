package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wczero/services/mp/api/internal/logic"
	"wczero/services/mp/api/internal/svc"
	"wczero/services/mp/api/internal/types"
)

func MPBasicSetupHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MPBasicSetupRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMPBasicSetupLogic(r.Context(), svcCtx)
		resp, err := l.MPBasicSetup(req)
		if err != nil || resp == nil {
			httpx.OkJson(w, "token illegal")
			return
		}
		httpx.OkJson(w, req.EchoStr)
	}
}
