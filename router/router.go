package router

import (
	"mvc_for_gin/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(router *gin.Engine) {
	router.GET("/", controller.Index)
	router.GET("/json", controller.JsonIndex)
	router.GET("/:username", controller.FindUser)
	router.POST("/add", controller.RegisterUser)
}
