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

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})

	v1 := r.Group("api/v1")
	{
		//用户操作
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authed := v1.Group("/")
		authed.Use(middleware.JWT()) // 使用token鉴权中间件
		{
			authed.POST("task", api.CreateTask)  // 创建备忘录
			authed.GET("task/:id", api.ShowTask) // 展示备忘录
		}
	}

	return r
}
