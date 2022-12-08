package middleware

import (
	"code/utils"
	"net/http"
	"strings"
)

type MiddlewareMiddleware struct {
}

func NewMiddlewareMiddleware() *MiddlewareMiddleware {
	return &MiddlewareMiddleware{}
}

func (m *MiddlewareMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "Bearer" || token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token为空！！！"))
			return
		}
		tokenlist := strings.SplitN(token, " ", 2)
		if tokenlist[0] != "Bearer" && len(tokenlist) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token格式错误！！！"))
			return
		}
		// 验证 Token
		userClaim, err := utils.DecodeToken(tokenlist[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		r.Header.Set("UserIdentity", userClaim.Identity)
		next(w, r)
	}
}
