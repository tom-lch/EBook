package route

import (
	"admin/global"
	"admin/handle"
	"admin/middleware"
	"admin/tools/config"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.GinEngine == nil {
		r = gin.Default()
	} else {
		r = global.GinEngine
	}
	if config.SslConfig.Enable {
		r.Use(handle.TlsHandler())
	}
	middleware.InitMiddleware(r)
	authMiddleware, err := middleware.AuthInit()
	tools.HasError(err, "JWT Init Error", 5000)
	InitSysRouter(r, authMiddleware)
	return r
}
