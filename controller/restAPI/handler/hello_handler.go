package handler

import (
	"github.com/ccchieh/ginHelper"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 第一个API
// @Description 这是第一个API
// @Tags 文章
// @Success 200
// @Failure 400
// @Router /HelloHandler [get]
// @version 1.0
func (h *Helper) HelloHandler() (r *ginHelper.Router) {
	return &ginHelper.Router{
		Path:   "/HelloHandler",
		Method: "GET",
		Handlers: []gin.HandlerFunc{
			helloHandler,
		}}
}
func helloHandler(c *gin.Context) {

	c.String(http.StatusOK, "Hello world!")
}
