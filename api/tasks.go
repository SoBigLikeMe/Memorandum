package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"memorandum/pkg/utils"
	"memorandum/service"
)

// CreateTask ShowAccount godoc
// @Summary      创建新的备忘录
// @Description  传入token解析id绑定到备忘录数据
// @Tags         accounts
// @Accept       application/json
// @Success      200 {object}
// @Router       /api/v1/CreateTask
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

// ShowTask
// @Summary		 创建新的备忘录
// @Description  传入token从请求头中获取备忘录id返回备忘录(json)
// @Tags		 tid
// @Accept       application/json
// @Success      200 {object}
// @Router       /api/v1/ShowTask/:id
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

// ListTasks
// @Summary		 展示所有备忘录
// @Description  解析token拿到id，查询id所创建的备忘录
// @Tags		 id
// @Accept       application/json
// @Success      200 {object}
// @Router       /api/v1/ListTasks
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

// UpdateTask
// @Summary		 更新备忘录
// @Description  解析token拿到id，查询id更新的备忘录
// @Tags		 tid
// @Accept       application/json
// @Success      200 {object}
// @Router       /api/v1/UpdateTask
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

// SearchTask
// @Summary		 模糊查询备忘录
// @Description
// @Tags		 title
// @Accept       application/json
// @Success      200 {object}
// @Router       /api/v1/SearchTask
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

// DeleteTask
// @Summary		 删除备忘录
// @Description	 解析token,拿到id删除对应的备忘录
// @Tags		 id
// @Accept       application/json
// @Success      200 {object}
// @Router       /api/v1/SearchTask
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
