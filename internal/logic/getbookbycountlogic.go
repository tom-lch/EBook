package logic

import (
	"EBook/internal/cudr"
	"context"

	"EBook/internal/svc"
	"EBook/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBookByCountLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookByCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBookByCountLogic {
	return GetBookByCountLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBookByCountLogic) GetBookByCount(req types.BookCountReq) (*types.BookCountResp, error) {
	// todo: add your logic here and delete this line
	// 根据Count获取book list
	count := req.Count
	lists, err := cudr.GetCountBooks(count, l.svcCtx.DB)
	if err != nil {
		return nil, err
	}
	return &types.BookCountResp{Lists: lists}, nil
}
