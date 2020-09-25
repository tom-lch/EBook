package logic

import (
	"context"

	"EBook/internal/svc"
	"EBook/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserInfoLogic {
	return GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req types.GetUserInfoReq) (*types.GetUserInfoResp, error) {

	return &types.GetUserInfoResp{NickName: "dfdf", Email: "wef"}, nil
}
