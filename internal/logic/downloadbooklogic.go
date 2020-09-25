package logic

import (
	"context"

	"EBook/internal/svc"
	"EBook/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DownloadBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownloadBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) DownloadBookLogic {
	return DownloadBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DownloadBookLogic) DownloadBook(req types.BookDetailReq) (*types.BookUploadResp, error) {
	// todo: add your logic here and delete this line

	return &types.BookUploadResp{}, nil
}
