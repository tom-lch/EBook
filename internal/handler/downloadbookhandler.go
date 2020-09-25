package handler

import (
	"net/http"

	"EBook/internal/logic"
	"EBook/internal/svc"
	"EBook/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func downloadBookHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BookDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewDownloadBookLogic(r.Context(), ctx)
		resp, err := l.DownloadBook(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.WriteJson(w, http.StatusOK, resp)
		}
	}
}
