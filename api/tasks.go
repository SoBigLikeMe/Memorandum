package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"memorandum/pkg/utils"
	"memorandum/service"
)

// CreateTask 创建备忘录
func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization")) //传入token
	if err := c.ShouldBind(&createTask); err == nil {          //绑定前端传入的数据
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Println(err)
		c.JSON(400, err)
	}
}

// ShowTask 展示一条备忘录
func ShowTask(c *gin.Context) {
	var ShowTask service.ShowTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&ShowTask); err == nil {
		res := ShowTask.Show(c.Param("id"), claim.Id) // 从请求头中拿到id
		c.JSON(200, res)
	} else {
		logging.Println(err)
		c.JSON(400, err)
	}
}

// ListTasks 展示所有备忘录
func ListTasks(c *gin.Context) {
	listService := service.ListTasksService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error()
		c.JSON(400, err)
	}
}

// UpdateTask 更新备忘录
func UpdateTask(c *gin.Context) {
	updateTaskService := service.UpdateTaskService{}
	if err := c.ShouldBind(&updateTaskService); err == nil {
		res := updateTaskService.Update(c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error()
		c.JSON(400, err)
	}
}

// SearchTask 模糊查找备忘录
func SearchTask(c *gin.Context) {
	searchTaskService := service.SearchTaskService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&searchTaskService); err == nil {
		res := searchTaskService.Search(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Error()
		c.JSON(400, err)
	}
}

// DeleteTask 删除备忘录
func DeleteTask(c *gin.Context) {
	deleteTaskService := service.DeleteTaskService{}
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteTaskService); err == nil {
		res := deleteTaskService.Delete(claim.Id, c.Param("id"))
		c.JSON(200, res)
	} else {
		logging.Error()
		c.JSON(400, err)
	}
}
