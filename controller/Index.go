package controller

import (
	"mvc_for_gin/models"
	"mvc_for_gin/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello World",
	})
}

func Look(c *gin.Context) {
	users, err := service.FindAll()
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "添加失败",
		})
	} else {
		c.JSON(200, gin.H{
			"msg":  "添加成功",
			"data": users,
		})
	}
}

func AddUser(c *gin.Context) {
	name := c.Query("name")

	user := models.User{Name: name, Age: 18, Level: 6}
	err := service.AddUser(&user)
	if err != nil {
		c.JSON(200, gin.H{
			"msg": "添加失败",
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "添加成功",
		})
	}
}

func DelUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	rows := service.DelUser(id)
	if rows == 0 {
		c.JSON(200, gin.H{
			"msg": "删除失败",
		})
	} else {
		c.JSON(200, gin.H{
			"msg": "删除成功",
		})
	}
}
