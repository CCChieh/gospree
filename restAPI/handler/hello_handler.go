package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Test struct {
	Name string
}

// @Summary 第一个API
// @Description 实际并没有什么用只是用来看看的
// @Tags 文章
// @Success 200 object Test 成功后返回值
// @Failure 400
// @Router /HelloHandler [get]
// @version 1.0
func HelloHandler(c *gin.Context) {

	c.JSON(http.StatusOK, Test{Name: "Zheng"})
}
