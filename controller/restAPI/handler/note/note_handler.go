package note

import (
	"github.com/ccchieh/ginHelper"
	"github.com/ccchieh/gospree/controller/restAPI/middleware"
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/model"
	"github.com/ccchieh/gospree/service"
	"github.com/ccchieh/gospree/util/ret"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Tags Note
// @Summary 新建Note
// @Param note body service.CreateNoteParams true "Note"
// @Success 200
// @Failure 400
// @Router /note [post]
// @version 1.0
func (h *Helper) CreateNoteHandler() (r *ginHelper.Router) {

	handler := func(c *gin.Context) {
		params := new(service.CreateNoteParams)
		if err := c.ShouldBind(params); err != nil {
			core.Log.Info(err)
			ret.Result(c, http.StatusBadRequest, nil, core.ErrParameterMatch)
			return
		}
		value, exists := c.Get("user")
		if !exists {
			ret.Result(c, http.StatusBadRequest, nil, core.ErrPermission)
			return
		}
		user := value.(*model.User)
		params.ID = user.ID
		note, err := service.CreateNoteService(params)
		if err != nil {
			ret.Result(c, http.StatusBadRequest, nil, err)
			return
		}
		ret.Result(c, http.StatusOK, gin.H{"noteID": note.ID}, nil)
	}

	return &ginHelper.Router{
		Path:   "/",
		Method: "POST",
		Handlers: []gin.HandlerFunc{
			middleware.TokenMiddleware(),
			handler,
		}}
}

// @Tags Note
// @Summary 获取Note
// @Param noteID query int true "Note"
// @Success 200
// @Failure 400
// @Router /note [get]
// @version 1.0
func (h *Helper) GetNoteHandler() (r *ginHelper.Router) {

	handler := func(c *gin.Context) {
		params := new(service.GetNoteParams)
		if err := c.ShouldBind(params); err != nil {
			core.Log.Info(err)
			ret.Result(c, http.StatusBadRequest, nil, core.ErrParameterMatch)
			return
		}

		note, err := service.GetNoteService(params)
		if err != nil {
			ret.Result(c, http.StatusBadRequest, nil, err)
			return
		}
		ret.Result(c, http.StatusOK, gin.H{"content": note.Content}, nil)
	}

	return &ginHelper.Router{
		Path:   "/",
		Method: "GET",
		Handlers: []gin.HandlerFunc{
			handler,
		}}
}

// @Tags Note
// @Summary 获取NoteList
// @Param page query int true "Note"
// @Success 200
// @Failure 400
// @Router /note/list [get]
// @version 1.0
func (h *Helper) GetNoteIDListHandles() (r *ginHelper.Router) {

	handler := func(c *gin.Context) {
		params := new(service.GetNoteIDListParams)
		if err := c.ShouldBind(params); err != nil {
			core.Log.Info(err)
			ret.Result(c, http.StatusBadRequest, nil, core.ErrParameterMatch)
			return
		}

		noteIDs, err := service.GetNoteIDListService(params)
		if err != nil {
			ret.Result(c, http.StatusBadRequest, nil, err)
			return
		}
		ret.Result(c, http.StatusOK, gin.H{"noteIDList": noteIDs}, nil)
	}

	return &ginHelper.Router{
		Path:   "/list",
		Method: "GET",
		Handlers: []gin.HandlerFunc{
			handler,
		}}
}
