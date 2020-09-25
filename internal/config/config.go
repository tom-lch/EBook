package config

import (
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DBs DataBase
	JA  JwtAuth
	Ct  Captcha
}

type Captcha struct {
	CaptchaImgHeight int
	CaptchaImgWidth  int
	CaptchaKeyLong   int
}

type DataBase struct {
	DBdrive  string
	DBsource string
}

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}
