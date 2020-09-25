package logic

import (
	"EBook/internal/cudr"
	"EBook/internal/model"
	"EBook/internal/svc"
	"EBook/internal/types"
	"EBook/pkg/utils"
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
	"time"
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
	jwtTokenResp, err := l.JWT(req)
	if err != nil {
		return nil, err
	}
	return &types.UserLoginResp{Username: u.Username, NickName: u.NickName, Token: token, JwtTokenResp: *jwtTokenResp}, nil
	// return nil, errors.New("验证码错误")
}

func (l *UserLoginLogic) JWT(req types.UserLoginReq) (*types.JwtTokenResp, error) {
	var accessExpire = l.svcCtx.Config.JA.AccessExpire
	var accessSecret = l.svcCtx.Config.JA.AccessSecret
	now := time.Now().Unix()
	payloads := map[string]interface{}{
		"iss": req.Username,
		"sub": "login",
	}
	accessToken, err := l.genToken(now, accessSecret, payloads, accessExpire)
	if err != nil {
		logx.Error("生成token出错", err)
		return nil, err
	}
	// 将刷新时间设置为过期直接的一半
	return &types.JwtTokenResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *UserLoginLogic) genToken(iat int64, secret string, payloads map[string]interface{}, seconds int64) (string, error) {
	// 生成token遵循jwt的生成规则，将 header、payload使用base64进行编码，然后用secert经SHA256分别生成token
	// 首先获取到claims
	Claims := make(jwt.MapClaims)
	Claims["exp"] = iat + seconds
	Claims["iat"] = iat
	for k, v := range payloads {
		Claims[k] = v
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = Claims
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		logx.Error("token.SignedString 出错", err, secret)
		return "", err
	}
	return tokenString, nil
}
