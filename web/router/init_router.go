package router

import (
	"github.com/gin-gonic/gin"
	"mail.web/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())
	SetupApiRouters(r)
	return r
}
