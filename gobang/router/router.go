package router

import (
	"github.com/gin-gonic/gin"
	"gobang/websocket"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	m := websocket.InitMelody()
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})
	return r
}