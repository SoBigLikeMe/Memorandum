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
