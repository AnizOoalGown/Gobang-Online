package main

import (
	"gobang/config"
	"gobang/router"
	"strconv"
)

func main() {
	addr := ":" + strconv.Itoa(config.Config.Get("server.port").(int))
	r := router.InitRouter()
	r.Run(addr)
}
