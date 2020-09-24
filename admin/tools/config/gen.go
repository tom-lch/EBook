package config

import "github.com/spf13/viper"

type Gen struct {
	Dbname string
	Frontpath string
}


func InitGen(cfg *viper.Viper) *Gen {
	return &Gen{
		Dbname: cfg.GetString("dbname"),
		Frontpath: cfg.GetString("frontpath"),
	}
}

var GenConfig = new(Gen)