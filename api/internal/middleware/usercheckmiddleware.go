package middleware

import (
	"STFrontground-backend/rpc/pkg/errcode"
	"encoding/json"
	"fmt"
	"github.com/tal-tech/go-zero/rest/httpx"
	"net/http"
)

const (
	userKey = `x-user-id`
)

type UsercheckMiddleware struct {
}

func NewUsercheckMiddleware() *UsercheckMiddleware {
	return &UsercheckMiddleware{}
}

func (m *UsercheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get(userKey)
		jwtUserId := r.Context().Value("userId")
		userInt, err := json.Number(userId).Int64()
		if err != nil {
			httpx.Error(w, errcode.ErrorUserInfo)
			return
		}

		jwtInt, err := json.Number(fmt.Sprintf("%v", jwtUserId)).Int64()
		if err != nil {
			httpx.Error(w, errcode.ErrorUserInfo)
			return
		}

		if jwtInt != userInt {
			httpx.Error(w, errcode.AuthDeny)
			return
		}

		next(w, r)
	}
}
