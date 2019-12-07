package ret

import (
	"github.com/ccchieh/gospree/core"
	"github.com/gin-gonic/gin"
	"time"
)

func Result(c *gin.Context, statusCode int, data interface{}, err error) {

	if err != nil {
		e, ok := err.(*retErr)
		if !ok {
			e = Unknown
			e.Code = 99999
			e.Message = "Unknown error: " + err.Error()
		}
		result := errorResult{baseResult{time.Now(), e.Code}, e.Message}
		c.JSON(statusCode, result)
		core.Log.Err(e)
	}
	if data != nil {
		result := okResult{baseResult{time.Now(), 0}, data}
		c.JSON(statusCode, result)
	}

}
