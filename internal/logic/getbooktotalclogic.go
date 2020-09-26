package logic

import (
	"EBook/internal/cudr"
	"context"

	"EBook/internal/svc"
	"EBook/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBookTotalCLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetBookTotalCLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetBookTotalCLogic {
	return GetBookTotalCLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetBookTotalCLogic) GetBookTotalC(req types.GetBookTotalcReq) (*types.GetBookTotalcResp, error) {
	// todo: add your logic here and delete this line
	// 直接进入数据库中获取到当前数据库的行数
	c, err := cudr.GetTotalC(l.svcCtx.DB)
	if err != nil {
		return nil, err
	}
	return &types.GetBookTotalcResp{Count: c}, nil
}
