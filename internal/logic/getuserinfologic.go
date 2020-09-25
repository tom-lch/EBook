package logic

import (
	"EBook/internal/cudr"
	"EBook/internal/model"
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
	// 在jwt认证的情况下，使用用户访问数据库获取用户信息
	u := &model.User{Username: req.Username}
	user, err := cudr.GetUserInfo(u, l.svcCtx.DB)
	if err != nil {
		return nil, err
	}
	return &types.GetUserInfoResp{NickName: user.Username + user.NickName, Email: user.Email}, nil
}
