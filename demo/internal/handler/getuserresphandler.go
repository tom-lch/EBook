package handler

import (
	"net/http"

	"demo/internal/logic"
	"demo/internal/svc"
	"demo/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func getUserRespHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserResp
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetUserRespLogic(r.Context(), ctx)
		resp, err := l.GetUserResp(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
