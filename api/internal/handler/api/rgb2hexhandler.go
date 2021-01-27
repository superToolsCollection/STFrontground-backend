package handler

import (
	"net/http"

	"STFrontground-backend/api/internal/logic/api"
	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func Rgb2hexHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Rgb2HexReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewRgb2hexLogic(r.Context(), ctx)
		resp, err := l.Rgb2hex(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
