package note

import (
	"github.com/ccchieh/ginHelper"
	"github.com/ccchieh/gospree/controller/restAPI/middleware"
	"github.com/ccchieh/gospree/service"
	"github.com/gin-gonic/gin"
)

// @Tags Note
// @Summary 新建Note
// @Param note body service.CreateNote true "Note"
// @Success 200
// @Failure 400
// @Router /note [post]
// @version 1.0
func (h *Helper) CreateNoteHandler() (r *ginHelper.Router) {

	return &ginHelper.Router{
		Param:  new(service.CreateNote),
		Path:   "/",
		Method: "POST",
		Handlers: []gin.HandlerFunc{
			middleware.TokenMiddleware(),
		},
	}
}

// @Tags Note
// @Summary 获取Note
// @Param noteID query int true "Note"
// @Success 200
// @Failure 400
// @Router /note [get]
// @version 1.0
func (h *Helper) GetNoteHandler() (r *ginHelper.Router) {
	return &ginHelper.Router{
		Param:  new(service.GetNote),
		Path:   "/",
		Method: "GET",
	}
}

// @Tags Note
// @Summary 获取NoteList
// @Param page query int true "Note"
// @Success 200
// @Failure 400
// @Router /note/list [get]
// @version 1.0
func (h *Helper) GetNoteListHandles() (r *ginHelper.Router) {
	return &ginHelper.Router{
		Param:  new(service.GetNoteList),
		Path:   "/list",
		Method: "GET",
	}
}
