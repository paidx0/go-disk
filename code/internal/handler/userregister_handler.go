package handler

import (
	"net/http"

	"code/internal/logic"
	"code/internal/svc"
	"code/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func userregisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegisterRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserregisterLogic(r.Context(), svcCtx)
		resp, err := l.Userregister(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
