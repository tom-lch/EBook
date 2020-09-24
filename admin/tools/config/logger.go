package config

import "github.com/spf13/viper"

type Logger struct {
	// 日志存放路径
	Path string
	// 控制台日志
	Stdout bool
	// 日志等级
	Level string
	// 业务日志开关
	Enabledbus bool
	// 请求日志开关
	Enabledreq bool
	// 数据库日志开关 dev模式，将自动开启
	Enableddb bool
	// 自动任务日志开关 dev模式，将自动开启
	Enabledjob bool
}


func InitLogger(cfg *viper.Viper) *Logger {
	return &Logger{
		Path:   cfg.GetString("path"),
		Stdout: cfg.GetBool("stdout"),
		Level:  cfg.GetString("level"),
		Enabledbus: cfg.GetBool("enabledbus"),
		Enabledreq: cfg.GetBool("enabledreq"),
		Enableddb:  cfg.GetBool("enableddb"),
		Enabledjob: cfg.GetBool("enabledjob"),
	}
}

var LoggerConfig = new(Logger)
