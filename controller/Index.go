package controller

import (
	"mvc_for_gin/libs/response"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// 注意，此处的code并不是http状态码而是业务上的状态码
	response.Msg(c, 200, "Hello MVC_For_Gin")
}
