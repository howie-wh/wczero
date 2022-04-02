package handler

import (
	"net/http"
	"wczero/common/response" //1

	"github.com/zeromicro/go-zero/rest/httpx"
	"wczero/services/wallpaper/api/internal/logic"
	"wczero/services/wallpaper/api/internal/svc"
	"wczero/services/wallpaper/api/internal/types"
)

func ImportHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImportRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewImportLogic(r.Context(), svcCtx)
		resp, err := l.Import(req)
		response.Response(w, resp, err) //2
	}
}
