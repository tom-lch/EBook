package config

import (
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DBdrive          string
	DBsource         string
	CaptchaImgHeight int
	CaptchaImgWidth  int
	CaptchaKeyLong   int
}
