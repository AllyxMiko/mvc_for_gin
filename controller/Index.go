package controller

import (
	"mvc_for_gin/libs/response"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	response.StringMsg(c, "Hello MVC_For_Gin")
}

func JsonIndex(c *gin.Context) {
	response.JsonMsg(c, 0, "Hello MVC_For_Gin")
}
