// @title           备忘录
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

package main

import (
	"memorandum/config"
	"memorandum/model"
	"memorandum/router"
)

func main() {
	defer func() {
		err := model.DB.Close()
		if err != nil {
			return
		}
	}()
	config.Init()
	r := router.NewRouter()
	_ = r.Run(config.HttpPort)
}
