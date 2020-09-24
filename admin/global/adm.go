package global

import (
	"github.com/gin-gonic/gin"
	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
	"github.com/robfig/cron/v3"
)



var GinEngine *gin.Engine
var CasbinEnforcer *casbin.SyncedEnforcer
var GDB *gorm.DB

var GADMCron *cron.Cron


var (
	Source string
	Driver string
	DBName string
)

var Version string

func init() {
	Version = "0.0.1"
}

