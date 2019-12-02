package middleware

import (
	"fmt"
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/util"
	"github.com/gin-gonic/gin"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		core.Log.Info(fmt.Sprintf("%s %s %s %s %d",
			util.RemoteIp(c.Request),
			c.Request.Method,
			c.Request.RequestURI,
			time.Since(start),
			c.Writer.Status(),
		))
	}
}
