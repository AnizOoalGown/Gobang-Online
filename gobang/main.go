package main

import (
	"gobang/config"
	"gobang/router"
)

func main() {
	addr := ":" + config.Config.Get("server.port").(string)
	r := router.InitRouter()
	r.Run(addr)
}
