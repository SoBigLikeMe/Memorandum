package main

import (
	"memorandum/config"
	"memorandum/router"
)

func main() {
	config.Init()
	r := router.NewRouter()
	_ = r.Run(config.HttpPort)
}
