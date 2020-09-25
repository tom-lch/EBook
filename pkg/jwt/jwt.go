package jwtAuth

import (
	"EBook/internal/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
	"time"
)

func JWT(AccessExpire int64, AccessSecret string) (*types.JwtTokenResp, error) {
	var accessExpire = AccessExpire
	var accessSecret = AccessSecret
	now := time.Now().Unix()
	payloads := map[string]interface{}{
		"sub": "login",
	}
	accessToken, err := genToken(now, accessSecret, payloads, accessExpire)
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

func genToken(iat int64, secret string, payloads map[string]interface{}, seconds int64) (string, error) {
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
