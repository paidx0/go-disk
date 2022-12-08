package handler

import (
	"net/http"

	"code/internal/logic"
	"code/internal/svc"
	"code/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userdetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserdetailLogic(r.Context(), svcCtx)
		resp, err := l.Userdetail(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
