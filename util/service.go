package util

import (
	"github.com/ccchieh/gospree/model"
	"github.com/gin-gonic/gin"
)

func GetIDFromContext(c *gin.Context) uint {
	value, exists := c.Get("user")
	if !exists {
		return 0
	}
	user := value.(*model.User)
	return user.ID
}
