package libs

import (
	"mvc_for_gin/configs"
	db "mvc_for_gin/database"
	"mvc_for_gin/router"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Port       int
	Http       *gin.Engine
	noDataBase bool
}

var Server = new(HttpServer)

func (h *HttpServer) New() *HttpServer {
	h.Http = gin.New()
	return h
}

func (h *HttpServer) Default() *HttpServer {
	h.Http = gin.Default()
	return h
}

func (h *HttpServer) Use(middleware ...gin.HandlerFunc) *HttpServer {
	h.Http.Use(middleware...)
	return h
}

func (h *HttpServer) NoDataBase() *HttpServer {
	h.noDataBase = true
	return h
}

func (h *HttpServer) Run() {
	if !h.noDataBase {
		db.InitDataBase()
	}
	router.RegisterRouter(h.Http)
	h.Http.Run(configs.Host + ":" + configs.Port)
}
