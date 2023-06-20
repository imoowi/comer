/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package middlewares

import (
	"fmt"
	"{{.moduleName}}/utils"
	"math"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() // 调用该请求的剩余处理程序
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000))))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "Unknown"
		}
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		url := c.Request.RequestURI
		Log := utils.Logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"SpendTime": spendTime,
			"path":      url,
			"Method":    method,
			"status":    statusCode,
			"Ip":        clientIP,
			"DataSize":  dataSize,
			"UserAgent": userAgent,
			"context":   c,
		})

		if len(c.Errors) > 0 { // 内部错误
			go Log.Error(c.Errors.ByType(gin.ErrorTypePrivate))
		}
		if statusCode >= 500 {
			go Log.Error()
		} else if statusCode >= 400 {
			go Log.Warn()
		} else {
			go Log.Info()
		}
	}
}
