package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"moment/api-model/internal/logic"
	"moment/api-model/internal/svc"
	"moment/api-model/internal/types"
)

func get_tokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGet_tokenLogic(r.Context(), svcCtx)
		resp, err := l.Get_token(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
