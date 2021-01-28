package handler

import (
	"net/http"

	"STFrontground-backend/api/internal/logic/api"
	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func ToolsListByTagNameHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ToolsListByTagNameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewToolsListByTagNameLogic(r.Context(), ctx)
		resp, err := l.ToolsListByTagName(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
