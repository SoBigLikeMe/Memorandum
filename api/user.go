package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"memorandum/pkg/utils"
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
		fmt.Printf("%v\n", err)
		c.JSON(400, err)
	}
}

// ReturnID 返回用户id
func ReturnID(c *gin.Context) {
	var UserService service.UserService

	//绑定数据
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&UserService); err != nil {
		res := UserService.ReturnId(claim.UserName)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
