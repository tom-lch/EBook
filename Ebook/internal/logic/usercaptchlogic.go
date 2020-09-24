package logic

import (
	"context"

	"//internal/svc"
	"//internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type UserCaptchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCaptchLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserCaptchLogic {
	return UserCaptchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCaptchLogic) UserCaptch(req types.CaptchaReq) (*types.CaptchaResp, error) {
	// todo: add your logic here and delete this line

	return &types.CaptchaResp{}, nil
}
