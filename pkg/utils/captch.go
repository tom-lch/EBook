package utils

import (
	"github.com/mojocn/base64Captcha"
)

var Store = base64Captcha.DefaultMemStore

type SysCaptchaResponse struct {
	CaptchaId string `json:"captchaId"`
	PicPath   string `json:"picPath"`
}


