package ret

import (
	"github.com/ccchieh/gospree/core"
	"github.com/gin-gonic/gin"
	"time"
)

func Result(c *gin.Context, statusCode int, data interface{}, err error) {

	if err != nil {
		result := errorResult{baseResult{Time: time.Now()}, statusCode, err}
		c.JSON(statusCode, result)
		core.Log.Err(err)
		return
	}
	if data != nil {
		result := okResult{baseResult{Time: time.Now()}, data}
		c.JSON(statusCode, result)
	}

}
