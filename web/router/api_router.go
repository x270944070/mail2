package router

import (
	"github.com/gin-gonic/gin"
	"mail.web/controller"
)

func SetupApiRouters(r *gin.Engine) {
	v1 := r.Group("/api/v1")

	v1.POST("pipeline", controller.UserLogin)

}
