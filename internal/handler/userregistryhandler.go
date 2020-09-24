package handler

import (
	"github.com/tal-tech/go-zero/core/logx"
	"net/http"

	"EBook/internal/logic"
	"EBook/internal/svc"
	"EBook/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func userRegistryHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRegistryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserRegistryLogic(r.Context(), ctx)
		resp, err := l.UserRegistry(req)
		if err != nil {
			logx.Error(err)
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
