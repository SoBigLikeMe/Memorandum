package service

import (
	"github.com/jinzhu/gorm"
	logging "github.com/sirupsen/logrus"
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

type DeleteTaskService struct{}

type UpdateTaskService struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	Status  int    `json:"status" form:"status"` //0是未做，1是已做
}

type SearchTaskService struct {
	Info string `form:"info" json:"info"`
}

type ListTasksService struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start"`
}

// Create 创建一条备忘录
func (service *CreateTaskService) Create(id uint) serialzer.Response {
	var user model.User
	model.DB.First(&user.ID, id)
	task := model.Task{
		Model:     gorm.Model{},
		User:      user,
		Uid:       id,
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

// List 展示所有备忘录
func (service *ListTasksService) List(id uint) serialzer.Response {
	var tasks []model.Task
	var total int64
	if service.Limit == 0 {
		service.Limit = 15
	}
	model.DB.Model(model.Task{}).Preload("User").Where("uid = ?", id).Count(&total).
		Limit(service.Limit).Offset((service.Start - 1) * service.Limit).
		Find(&tasks)
	return serialzer.BuildListResponse(serialzer.BuildTasks(tasks), uint(total))
}

// Update 更新备忘录
func (service *UpdateTaskService) Update(id string) serialzer.Response {
	var task model.Task
	model.DB.Model(model.Task{}).Where("id = ?", id).First(&task)
	task.Content = service.Content
	task.Status = service.Status
	task.Title = service.Title
	code := 200
	err := model.DB.Save(&task).Error
	if err != nil {
		logging.Error()
		code = 400
		return serialzer.Response{
			Status: code,
			Msg:    "更新失败",
			Error:  err.Error(),
		}
	}
	return serialzer.Response{
		Status: code,
		Data:   "修改成功",
	}
}

// Search 模糊查找备忘录
func (s SearchTaskService) Search(id uint) serialzer.Response {
	var tasks []model.Task
	count := 0
	err := model.DB.Model(&model.Task{}).Preload("User").Where("uid=?", id).
		Where("title like ? or content like ?", "%"+s.Info+"%", "%"+s.Info+"%").
		Count(&count).Find(&tasks)
	if err == nil {
		return serialzer.Response{
			Status: 400,
			Data:   nil,
			Msg:    "查询失败",
		}
	}

	return serialzer.Response{
		Status: 200,
		Msg:    "查找成功",
		Data:   serialzer.BuildTasks(tasks),
	}
}

// Delete 删除备忘录
func (s DeleteTaskService) Delete(id uint, tid string) serialzer.Response {
	var task model.Task
	code := 200
	err := model.DB.Where("id = ? AND uid = ?", tid, id).Find(&task).Error

	if err != nil {
		code = 500
		return serialzer.Response{
			Status: code,
			Msg:    "权限不足",
		}
	}

	err = model.DB.Delete(&task).Error

	if err != nil {
		code = 200
		return serialzer.Response{
			Status: code,
			Msg:    "删除失败",
			Error:  err.Error(),
		}
	}
	return serialzer.Response{
		Status: code,
		Msg:    "删除成功",
	}
}
