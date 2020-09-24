package logger

import (
	"admin/global"
	"admin/tools"
	"admin/tools/config"
	"github.com/gogf/gf/os/glog"
)

var Logger *glog.Logger
var JobLogger *glog.Logger
var RequestLogger *glog.Logger

func Setup() {
	Logger = glog.New()
	_ = Logger.SetPath(config.LoggerConfig.Path+"/bus")
	Logger.SetStdoutPrint(config.LoggerConfig.Enabledbus && config.LoggerConfig.Stdout)
	Logger.SetFile("bus-{Y-m-d}.log")

	JobLogger = glog.New()
	_ = JobLogger.SetPath(config.LoggerConfig.Path+"/job")
	JobLogger.SetStdoutPrint(config.LoggerConfig.Enabledjob)
	JobLogger.SetFile("job-{Y-m-d}.log")

	RequestLogger = glog.New()
	_ = RequestLogger.SetPath(config.LoggerConfig.Path+"/req")
	RequestLogger.SetStdoutPrint(config.LoggerConfig.Enabledreq)
	RequestLogger.SetFile("req-{Y-m-d}.log")
	_ = RequestLogger.SetLevelStr(config.LoggerConfig.Level)

	Logger.Info(tools.Green("log init success!"))

	global.Logger = Logger
	global.JobLogger = JobLogger
	global.RequestLogger = RequestLogger
}
