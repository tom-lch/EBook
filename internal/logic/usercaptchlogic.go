package logic

import (
	"EBook/pkg/utils"
	"context"
	"github.com/mojocn/base64Captcha"

	"EBook/internal/svc"
	"EBook/internal/types"

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
	//字符,公式,验证码配置
	// 生成默认数字的driver
	driver := base64Captcha.NewDriverDigit(l.svcCtx.Config.Ct.CaptchaImgHeight, l.svcCtx.Config.Ct.CaptchaImgWidth, l.svcCtx.Config.Ct.CaptchaKeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, utils.Store)
	id, b64s, err := cp.Generate()
	if err != nil {
		logx.Error("验证码生成出错")
		return nil, err
	}
	return &types.CaptchaResp{CaptchaId: id, PicPath: b64s}, nil
}
