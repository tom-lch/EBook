package config

import "github.com/spf13/viper"

type Appliction struct {
	Mode 		  string
	Host 		  string
	Name 		  string
	Port 		  string
	JwtSecret     string
	DemoMsg       string
	Readtimeout   int
	Writertimeout int
	Enabledp 	  bool
}


func InitAppliction(cfg *viper.Viper) *Appliction {
	return &Appliction{
		Mode:cfg.GetString("mode"),
		Host: cfg.GetString("host"),
		Name: cfg.GetString("name"),
		Port: cfg.GetString("port"),
		JwtSecret: cfg.GetString("jwtSecret"),
		DemoMsg: cfg.GetString("demoMsg"),
		Readtimeout: cfg.GetInt("readtimeout"),
		Writertimeout: cfg.GetInt("writertimeout"),
		Enabledp: cfg.GetBool("enabledp"),
	}
}
var ApplictionConfig = new(Appliction)