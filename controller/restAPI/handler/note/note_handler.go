package note

import (
	"github.com/ccchieh/ginHelper"
	"github.com/ccchieh/gospree/controller/restAPI/middleware"
	"github.com/ccchieh/gospree/service"
	"github.com/gin-gonic/gin"
	"reflect"
)

// @Tags Note
// @Summary 新建Note
// @Param note body service.CreateNote true "Note"
// @Success 200
// @Failure 400
// @Router /note [post]
// @version 1.0
func (h *Helper) CreateNoteHandler() (r *ginHelper.Router) {
	handler := func(c *gin.Context) {
		params := new(service.CreateNote)
		params.ID = service.GetID(c)
		service.Process(c, params)
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

	// 取变量a的反射类型对象
	paramsType := reflect.TypeOf(service.GetNote{})
	// 根据反射类型对象创建类型实例

	handler := func(c *gin.Context) {
		params := reflect.New(paramsType).Interface().(service.Params)
		service.Process(c, params)
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
func (h *Helper) GetNoteListHandles() (r *ginHelper.Router) {
	handler := func(c *gin.Context) {
		params := new(service.GetNoteList)
		service.Process(c, params)
	}

	return &ginHelper.Router{
		Path:   "/list",
		Method: "GET",
		Handlers: []gin.HandlerFunc{
			handler,
		}}
}
