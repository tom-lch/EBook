package handler

import (
	"net/http"

	"demo/internal/logic"
	"demo/internal/svc"
	"demo/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func getInfoCallHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetInfoCallLogic(r.Context(), ctx)
		resp, err := l.GetInfoCall(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
