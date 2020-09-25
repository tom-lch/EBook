package logic

import (
	"EBook/internal/cudr"
	"EBook/internal/model"
	"EBook/internal/svc"
	"EBook/internal/types"
	jwtAuth "EBook/pkg/jwt"
	"EBook/pkg/utils"
	"context"
	"github.com/tal-tech/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req types.UserLoginReq) (*types.UserLoginResp, error) {
	// todo: add your logic here and delete this line
	// req中已经从web中获取到了绑定的json表单信息，下面要我们要验证该信息是否合法
	UserVerify := utils.Rules{
		"Username":  {utils.NotEmpty()},
		"NickName":  {utils.NotEmpty()},
		"Captcha":   {utils.NotEmpty()},
		"CaptchaId": {utils.NotEmpty()},
	}
	if UserVerifyErr := utils.Verify(req, UserVerify); UserVerifyErr != nil {
		logx.Error("该用户登录息有误, err:", UserVerifyErr)
		return nil, UserVerifyErr
	}
	// 判断验证码是否正确
	//if utils.Store.Verify(req.CaptchaId, req.Captcha, true) {
	//	// 下面去数据库中进行查找用户
	//	u := &model.User{Username: req.Username, Password: req.Password}
	//	u, err := cudr.Login(u, l.svcCtx.DB)
	//	if err != nil {
	//		return nil, err
	//	}
	//	// 生成Token
	//	token := utils.GenToken(u)
	//	return &types.UserLoginResp{Username: u.Username, NickName: u.NickName, Token: token}, nil
	//}
	// 验证码错误
	u := &model.User{Username: req.Username, Password: req.Password}
	u, err := cudr.Login(u, l.svcCtx.DB)
	if err != nil {
		return nil, err
	}
	token := utils.GenToken(u)
	// 在此处生成真正的Token
	jwtTokenResp, err := jwtAuth.JWT(l.svcCtx.Config.JA.AccessExpire, l.svcCtx.Config.JA.AccessSecret)
	if err != nil {
		return nil, err
	}
	return &types.UserLoginResp{Username: u.Username, NickName: u.NickName, Token: token, JwtTokenResp: *jwtTokenResp}, nil
	// return nil, errors.New("验证码错误")
}

