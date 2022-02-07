package middleware

import (
	"mvc_for_gin/database"

	"github.com/gin-gonic/gin"
)

func DbContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Db", database.DbConn)
		c.Next()
	}
}
