package service

import (
	"github.com/jinzhu/gorm"
	"memorandum/model"
	"memorandum/pkg/utils"
	"memorandum/serialzer"
)

type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=3,max=15" example:"FanOne"`
}

func (service UserService) Register() serialzer.Response {
	var user model.User
	var count int
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return serialzer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "该用户已存在",
			Error:  "",
		}
	}
	user.UserName = service.UserName
	//加密

	if err := user.SetPassword(service.Password); err != nil {
		return serialzer.Response{
			Status: 400,
			Data:   nil,
			Msg:    err.Error(),
			Error:  "",
		}
	}

	//创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serialzer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "数据库操作错误",
			Error:  "",
		}
	}

	return serialzer.Response{
		Status: 200,
		Data:   nil,
		Msg:    "创建成功",
		Error:  "",
	}
}

func (service UserService) Login() serialzer.Response {
	var user model.User

	//查询数据库中是否存在
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return serialzer.Response{
				Status: 400,
				Msg:    "用户不存在，请先登录",
			}
		}

		//其他因素导致无法登录
		return serialzer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "数据库错误",
			Error:  "",
		}
	}

	if user.CheckPassword(service.Password) == false {
		return serialzer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "密码错误",
			Error:  "",
		}
	}

	//发一个token，为了其他需要身份验证的功能给前端存储，例如创建备忘录

	token, err := utils.GenerateToken(user.ID, service.UserName, service.Password)
	if err != nil {
		return serialzer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "Token签发错误",
			Error:  "",
		}
	}
	return serialzer.Response{
		Status: 200,
		Data:   serialzer.TokenData{User: serialzer.BuildUser(user), Token: token},
		Msg:    "登陆成功",
	}
}
