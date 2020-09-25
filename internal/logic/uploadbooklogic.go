package logic

import (
	"EBook/internal/cudr"
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
	// 对书籍信息逻辑判断

	// 添加到数据库
	if err := cudr.CreateBook(req, l.svcCtx.DB); err != nil {
		return nil, err
	}
	resp := &types.BookUploadResp{
		BookId   : req.BookId,
		Bookname : req.Bookname,
		Status   : "OK",
		Code     : "1",
	}
	return resp, nil
}
