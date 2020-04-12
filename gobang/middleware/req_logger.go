package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/weekface/mgorus"
	"time"
)

func ReqLogger() gin.HandlerFunc {
	logger := logrus.New()
	hooker, err := mgorus.NewHooker("150.158.104.248:27017", "gobang", "log")
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
