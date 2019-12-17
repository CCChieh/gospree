package note

import (
	"github.com/ccchieh/ginHelper"
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/service"
	"github.com/ccchieh/gospree/util/ret"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags Note
// @Summary 新建Note
// @Param user body service.CreateNoteParams true "Note"
// @Success 200
// @Failure 400
// @Router /note [post]
// @version 1.0
func (h *Helper) CreateNoteHandler() (r *ginHelper.Router) {

	handler := func(c *gin.Context) {
		params := new(service.CreateNoteParams)
		if err := c.ShouldBind(params); err != nil {
			core.Log.Info(err)
			ret.Result(c, http.StatusBadRequest, nil, ret.ErrParameterMatch)
			return
		}

		_, err := service.CreateNoteService(params)
		if err != nil {
			ret.Result(c, http.StatusBadRequest, nil, err)
			return
		}
		ret.Result(c, http.StatusOK, nil, nil)
	}

	return &ginHelper.Router{
		Path:   "/",
		Method: "POST",
		Handlers: []gin.HandlerFunc{
			handler,
		}}
}
