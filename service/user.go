package service

import (
	"memorandum/model"
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
