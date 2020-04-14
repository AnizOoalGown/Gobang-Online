package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
	"gobang/config"
	"time"
)

func ReqLogger() gin.HandlerFunc {
	logger := logrus.New()
	addr := config.Config.Get("mongodb.addr").(string)
	db := config.Config.Get("mongodb.db").(string)
	collection := config.Config.Get("mongodb.collection").(string)
	hooker, err := mgorus.NewHooker(addr, db, collection)
	if err == nil {
		logger.Hooks.Add(hooker)
	}

	return func(c *gin.Context) {
		startTime := time.Now()
		reqProto := c.Request.Proto
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		fields := logrus.Fields{
			"protocol":    reqProto,
			"method":      reqMethod,
			"uri":         reqUri,
			"status_code": statusCode,
			"client_ip":   clientIP,
		}

		logger.WithFields(fields).Info("client connects")

		c.Next()

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime).Truncate(time.Second).Seconds()
		fields["latency_time"] = latencyTime

		logger.WithFields(fields).Info("client disconnects")
	}
}
