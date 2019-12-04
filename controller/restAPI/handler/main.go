package handler

import (
	"github.com/ccchieh/gospree/core"
	"github.com/gin-gonic/gin"
	"net/http"
)

type router struct {
	group    string
	path     string
	method   string
	handlers []gin.HandlerFunc
}

type helper struct {
}

// @Summary helper
// @Description 实际并没有什么用只是用来看看的
// @Tags 文章
// @Success 200 object Test 成功后返回值
// @Failure 400
// @Router /helper [get]
// @version 1.0
func (h *helper) HelperHandler() (r *router) {
	return &router{
		group:  "",
		path:   "/helper",
		method: "GET",
		handlers: []gin.HandlerFunc{
			helperHandler,
		}}
}

func helperHandler(c *gin.Context) {
	c.String(http.StatusOK, "Cool~")
}

func Build(r *gin.Engine) {
	h := new(helper)
	core.Log.Info(h)
}
