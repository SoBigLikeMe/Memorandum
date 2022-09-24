package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"memorandum/pkg/utils"
	"memorandum/service"
)

func CreateTask(c *gin.Context) {
	var createTask service.CreateTaskService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization")) //传入token
	if err := c.ShouldBind(&createTask); err == nil {
		res := createTask.Create(claim.Id)
		c.JSON(200, res)
	} else {
		logging.Println(err)
		c.JSON(400, err)
	}
}

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
