package middleware

import (
	"github.com/gin-gonic/gin"
	"memorandum/pkg/utils"
	"time"
)

//中间件进行token验证
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := 200 //返回状态
		token := c.GetHeader("Authorization")
		if token == "" {
			code = 400
		} else {
			claim, err := utils.ParseToken(token) //非空则解析token
			if err != nil {
				code = 404 //说明token无权限
			} else if time.Now().Unix() > claim.ExpiresAt {
				code = 401 //说明token超时，token无效
			}
		}

		if code != 200 {
			c.JSON(400, gin.H{
				"status": code,
				"msg":    "解析错误",
			})
			c.Abort() // 终止接下来的操作，终止接下来中间件的调用
			return
		}
		c.Next() //当处理完所有的中间件函数（包括本次请求）的时候才会停止，执行完一次完整的请求
	}
}
