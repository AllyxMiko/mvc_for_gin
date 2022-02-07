package server

import (
	db "mvc_for_gin/database"
	"mvc_for_gin/middleware"
	"mvc_for_gin/router"
	"mvc_for_gin/setting"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Http       *gin.Engine
	noDataBase bool
}

var Http = new(HttpServer)

func init() {
	setting.CheckConfigs()
	if !setting.PjtConfigs.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
}

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
		h.Use(middleware.DbContext())
	}
	h.Http.LoadHTMLGlob("views/*")
	router.RegisterRouter(h.Http)
	h.Http.Run(setting.PjtConfigs.Host + ":" + setting.PjtConfigs.Port)
}
