package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"moment/api-model/internal/logic"
	"moment/api-model/internal/svc"
	"moment/api-model/internal/types"
)

func del_momentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DelMomentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDel_momentLogic(r.Context(), svcCtx)
		resp, err := l.Del_moment(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
