package controller

import (
	"mvc_for_gin/libs/response"
	"mvc_for_gin/models"
	"mvc_for_gin/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 注册用户，为了演示方便，只要用户名，密码
func RegisterUser(c *gin.Context) {
	info := models.User{}
	c.BindJSON(&info)
	if info.Username == "" || info.Password == "" {
		response.JsonMsg(c, 500, "请求参数不完整！")
		return
	}
	// 查询用户是否已存在
	db, _ := c.Get("Db")
	if service.CheckUser(db.(*gorm.DB), info.Username) {
		response.JsonMsg(c, 400, "该用户已被注册！")
		return
	}
	// 注册用户
	if !service.AddUser(db.(*gorm.DB), &info) {
		response.JsonMsg(c, 500, "用户注册失败！")
		return
	}
	response.JsonMsg(c, 200, "注册成功！")
}

// 查询用户
func FindUser(c *gin.Context) {
	username := c.Param("username")
	db, _ := c.Get("Db")
	user := service.GetUser(db.(*gorm.DB), username)
	if user == nil {
		response.JsonMsg(c, 401, "该用户不存在！")
		return
	}
	response.SendData(c, 200, "查询成功！", user)
}
