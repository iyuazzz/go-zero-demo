package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"moment/api-model/internal/logic"
	"moment/api-model/internal/svc"
	"moment/api-model/internal/types"
)

func create_momentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateMomentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreate_momentLogic(r.Context(), svcCtx)
		resp, err := l.Create_moment(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
