package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"memorandum/api"
	"memorandum/middleware"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLGlob("template/*")

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"msg": "路径错误"})
	})

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})

	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register", api.UserRegister) //用户注册
		v1.POST("user/login", api.UserLogin)       //用户登录
		v1.GET("user/id", api.ReturnID)            //返回当前用户id
		authed := v1.Group("/")
		authed.Use(middleware.JWT()) // 使用token鉴权中间件
		{
			//备忘录操作
			authed.POST("task", api.CreateTask)     // 创建备忘录
			authed.GET("task/:id", api.ShowTask)    // 展示备忘录
			authed.GET("tasks", api.ListTasks)      //展示所有备忘录
			authed.PUT("task/:id", api.UpdateTask)  //更新备忘录
			authed.POST("search", api.SearchTask)   //查询备忘录
			authed.POST("task/:id", api.DeleteTask) //删除备忘录
		}
	}

	return r
}
