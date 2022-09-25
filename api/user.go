package api

import (
	"github.com/gin-gonic/gin"
	"memorandum/service"
)

// UserRegister 用户注册
func UserRegister(c *gin.Context) {
	var userRegister service.UserService

	//绑定注册数据
	if err := c.ShouldBind(&userRegister); err == nil {
		res := userRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	var userLogin service.UserService

	//绑定登录数据
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
