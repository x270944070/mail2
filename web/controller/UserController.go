package controller

import (
	"github.com/gin-gonic/gin"
	"log"
)

func UserLogin(ctx *gin.Context) {
	userName := ctx.GetString("userName")
	pwd := ctx.GetString("password")

	log.Info

}
