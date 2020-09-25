package logic

import (
	"EBook/internal/cudr"
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
	//根据page数和pagesize大小来获得指定位置的图书，
	page := req.Page
	pagesize := req.PageSize
	lists, err := cudr.GetPageSizeBooks(page, pagesize, l.svcCtx.DB)
	if err != nil {
		return nil, err
	}
	return &types.BookPagetResp{Lists: lists}, nil
}
