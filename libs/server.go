package libs

import (
	"mvc_for_gin/configs"
	db "mvc_for_gin/database"
	"mvc_for_gin/router"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Port        int
	Http        *gin.Engine
	useDataBase bool
}

var Server = new(HttpServer)

func (h *HttpServer) Default() *HttpServer {
	h.Http = gin.New()
	h.Http.Use(gin.Logger(), gin.Recovery())
	return h
}

func (h *HttpServer) UseDataBase() *HttpServer {
	h.useDataBase = true
	return h
}

func (h *HttpServer) Run() {
	if h.useDataBase {
		db.InitDataBase()
	}
	router.RegisterRouter(h.Http)
	h.Http.Run(":" + configs.Port)
}
