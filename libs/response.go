package libs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 封装了成功的响应和失败的响应，如需添加其他字段，请自行更改，此处仅作示例

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SuccessResponse struct {
	Response
	Data interface{} `json:"data"`
}

type WrongResponse struct {
	Response
}

func Msg(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
	})
}

func SendData(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Response{
			code,
			msg,
		},
		data,
	})
}

func Faild(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, WrongResponse{
		Response{
			code,
			msg,
		},
	})
}
