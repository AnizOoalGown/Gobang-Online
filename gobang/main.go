package main

import "gobang/router"

func main() {
	r := router.InitRouter()
	r.Run()
}