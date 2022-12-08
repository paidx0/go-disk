package handler

import (
	"net/http"

	"code/internal/logic"
	"code/internal/svc"
	"code/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func sharefilepullHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ShareFilePullRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSharefilepullLogic(r.Context(), svcCtx)
		resp, err := l.Sharefilepull(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
