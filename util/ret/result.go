package ret

import (
	"github.com/ccchieh/gospree/core"
	"github.com/gin-gonic/gin"
	"time"
)

func Result(c *gin.Context, statusCode int, data interface{}, err error) {

	if err != nil {
		e, ok := err.(*core.Err)
		if !ok {
			e = core.ErrUnknown
		}
		result := errorResult{baseResult{time.Now(), e.Code}, e.Message}
		c.JSON(statusCode, result)
		core.Log.Err(err)
	} else {
		result := okResult{baseResult{time.Now(), 0}, data}
		c.JSON(statusCode, result)
	}

}
