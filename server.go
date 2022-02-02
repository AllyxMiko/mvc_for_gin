package main

import (
	"log"
	db "mvc_for_gin/database"
	"mvc_for_gin/router"
	"mvc_for_gin/setting"
	"os"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Port       int
	Http       *gin.Engine
	noDataBase bool
}

func init() {
	_, err := os.Stat("./configs/config.yaml")
	if err != nil {
		_, err = os.Stat("./configs")
		if err != nil {
			os.Mkdir("./configs", 0777)
		}
		log.Println("配置文件不存在，系统将自动生成！请按照说明修改配置文件！")
		setting.CreateConfigsFile()
	}
	if !setting.PjtConfigs.Debug {
		gin.SetMode(gin.ReleaseMode)
	}
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
	h.Http.Run(setting.PjtConfigs.Host + ":" + setting.PjtConfigs.Port)
}
