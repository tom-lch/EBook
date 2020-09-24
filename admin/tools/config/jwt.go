package config

import (
	"github.com/spf13/viper"
)

type JWT struct {
	Secret string
	Timeout int64
}

func InitJwt(cfg *viper.Viper) *JWT {
	return &JWT{
		Secret: cfg.GetString("secret"),
		Timeout: cfg.GetInt64("timeout"),
	}
}


var JwtConfig = new(JWT)