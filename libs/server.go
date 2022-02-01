package libs

import (
	"mvc_for_gin/configs"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Port int
	Http *gin.Engine
}

var Server = new(HttpServer)

func (h *HttpServer) Default() *HttpServer {
	h.Http = gin.New()
	h.Http.Use(gin.Logger(), gin.Recovery())
	return h
}

func (h *HttpServer) Run() {
	h.Http.Run(":" + configs.Port)
}

// func InitServer() (http *gin.Engine) {
// 	// 初始化数据库
// 	// db.InitDataBase()

// 	// 启动http服务
// http = gin.New()
// 	// 应用全局中间件, logger为日志，recovery为错误恢复
// 	http.Use(gin.Logger(), gin.Recovery())
// 	// 注册路由
// 	router.RegisterRouter(http)

// 	return
// }
