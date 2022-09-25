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

// ShowTask 展示备忘录
func ShowTask(c *gin.Context) {
	var ShowTask service.ShowTaskService
	if err := c.ShouldBind(&ShowTask); err == nil {
		res := ShowTask.Show(c.Param("id")) // 从请求头中拿到id
		c.JSON(200, res)
	} else {
		logging.Println(err)
		c.JSON(400, err)
	}
}

// ListTasks 展示所有备忘录
func ListTasks(c *gin.Context) {
	listService := service.ListTasksService{}
	chaim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&listService); err == nil {
		res := listService.List(chaim.Id)
		c.JSON(200, res)
	} else {
		logging.Error()
		c.JSON(400, err)
	}
}

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
