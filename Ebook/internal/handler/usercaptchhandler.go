package handler

import (
	"net/http"

	"//internal/logic"
	"//internal/svc"
	"//internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func userCaptchHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserCaptchLogic(r.Context(), ctx)
		resp, err := l.UserCaptch(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
