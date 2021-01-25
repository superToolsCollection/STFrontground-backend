package handler

// 用户信息通过请求头中通过x-user-id携带到服务器
import (
	"net/http"

	"STFrontground-backend/api/internal/logic"
	"STFrontground-backend/api/internal/svc"
	"github.com/tal-tech/go-zero/rest/httpx"
)

func userInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("x-user-id")
		l := logic.NewUserInfoLogic(r.Context(), ctx)
		resp, err := l.UserInfo(userId)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
