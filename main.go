package main

import (
	db "mvc_for_gin/database"
	"mvc_for_gin/router"

	"github.com/gin-gonic/gin"
)

func init() {
	// 初始化数据库
	db.InitDataBase()
}

func main() {
	// http服务
	httpServer := gin.Default()
	// 注册路由
	router.RegisterRouter(httpServer)
	// 运行
	httpServer.Run()
}
