package logic

import (
	"context"

	"EBook/internal/svc"
	"EBook/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBookByPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookByPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBookByPageLogic {
	return GetBookByPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBookByPageLogic) GetBookByPage(req types.BookPagetReq) (*types.BookPagetResp, error) {
	// todo: add your logic here and delete this line

	return &types.BookPagetResp{}, nil
}
