package logic

import (
	"context"

	"EBook/internal/svc"
	"EBook/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UploadBookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) UploadBookLogic {
	return UploadBookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadBookLogic) UploadBook(req types.BookUploadReq) (*types.BookUploadResp, error) {
	// todo: add your logic here and delete this line

	return &types.BookUploadResp{}, nil
}
