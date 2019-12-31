package service

import (
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/model"
	"github.com/ccchieh/gospree/util"
	"github.com/gin-gonic/gin"
)

type CreateNote struct {
	Param
	Title   string `json:"title" binding:"required"`
	PreView string `json:"preView" binding:"required"`
	Content string `json:"content" binding:"required"`
	ID      uint   `json:"-"`
}

func (param *CreateNote) BeforeBind(c *gin.Context) {
	param.ID = util.GetIDFromContext(c)
	if param.ID == 0 {
		param.Err = core.ErrPermission
	}
}

func (param *CreateNote) Service() {
	note := &model.Note{
		Title:    param.Title,
		Content:  param.Content,
		PreView:  param.PreView,
		AuthorID: param.ID,
	}
	param.Err = note.CreateNote()
	param.Ret = gin.H{"noteID": note.ID}
	return
}

type GetNote struct {
	Param
	NoteID uint `form:"noteID" binding:"required"`
}

func (param *GetNote) Service() {
	note := new(model.Note)
	note.ID = param.NoteID
	core.Log.Info(param.NoteID)
	if err := note.GetNote(); err != nil {
		param.Err = core.ErrNoteNotFound
		return
	}
	param.Ret = gin.H{"title": note.Title, "preView": note.PreView, "content": note.Content}
}

type GetNoteList struct {
	Param
	Page int `form:"page" binding:"required"`
}

func (param *GetNoteList) Service() {
	notes := new(model.Notes)
	noteList, err := notes.GetNoteIDList(param.Page, core.Conf.GetNoteNumOfPage())
	if err != nil {
		param.Err = err
		return
	}
	if len(noteList) == 0 {
		param.Err = core.ErrEndOfNoteList
		return
	}
	param.Ret = gin.H{"noteList": noteList}
}
