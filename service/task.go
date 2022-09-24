package service

import (
	"github.com/jinzhu/gorm"
	"memorandum/model"
	"memorandum/serialzer"
	"time"
)

type CreateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做，1是已做
}

type ShowTaskService struct{}

// Create 创建一条备忘录
func (service *CreateTaskService) Create(id uint) serialzer.Response {
	var user model.User
	model.DB.First(&user.ID, id)
	task := model.Task{
		Model:     gorm.Model{},
		User:      user,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StratTime: time.Now().Unix(),
		EndTime:   0,
	}

	err := model.DB.Create(&task).Error
	if err != nil {
		return serialzer.Response{
			Status: 500,
			Data:   nil,
			Msg:    "创建备忘录失败",
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

//Show 展示一条备忘录
func (s ShowTaskService) Show(tid string) serialzer.Response {
	var task model.Task
	code := 200
	err := model.DB.Find(&task, tid).Error
	if err != nil {
		code = 500
		return serialzer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serialzer.Response{
		Status: code,
		Data:   serialzer.BuildTask(task),
		Msg:    "",
		Error:  "",
	}
}
