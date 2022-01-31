package middleware

import "github.com/gin-gonic/gin"

// 中间件编写示例

func ExceptFavicon() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Param("surl") == "favicon.ico" {
			c.Abort()
		}
		c.Next()
	}
}
