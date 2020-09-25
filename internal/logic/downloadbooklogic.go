package logic

import (
	"EBook/internal/cudr"
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

func (l *DownloadBookLogic) DownloadBook(req types.BookDetailReq) (*types.BookDetailResp, error) {
	// todo: add your logic here and delete this line
	// 根据书名和id获取下载链接
	bookname := req.Bookname
	bookid := req.BookId
	book, err :=cudr.DownloadBookByID(bookname, bookid, l.svcCtx.DB)
	if err != nil {
		return nil, err
	}
	resp := &types.BookDetailResp{
		Bookname: book.Bookname,
		BookId: book.ID,
		Author: book.Author,
		Uploader: book.Uploader,
		Cover: book.Cover,
		StoreAddr: book.StoreAddr,
	}
	return resp, nil
}
