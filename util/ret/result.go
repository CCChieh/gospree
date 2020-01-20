package ret

import (
	"github.com/ccchieh/gospree/core"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Result(c *gin.Context, data interface{}, err error) {
	if err != nil {
		e, ok := err.(*core.Err)
		if !ok {
			e = core.ErrUnknown
		}
		result := errorResult{baseResult{time.Now().Unix(), e.Code}, e.Message}
		c.JSON(e.HttpStatusCode, result)
		core.Log.Err(err)
	} else {
		result := okResult{baseResult{time.Now().Unix(), 0}, data}
		c.JSON(http.StatusOK, result)
	}
}
