package service

import (
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/model"
	"github.com/gin-gonic/gin"
)

type CreateNote struct {
	baseParams `json:"-"`
	Title      string `json:"title" binding:"required"`
	PreView    string `json:"preView" binding:"required"`
	Content    string `json:"content" binding:"required"`
	ID         uint   `json:"-"`
}

func (params *CreateNote) Service() {
	note := &model.Note{
		Title:    params.Title,
		Content:  params.Content,
		PreView:  params.PreView,
		AuthorID: params.ID,
	}
	params.err = note.CreateNote()
	params.result = gin.H{"noteID": note.ID}
	return
}

type GetNote struct {
	baseParams `json:"-"`
	NoteID     uint `form:"noteID" binding:"required"`
}

func (params *GetNote) Service() {
	note := new(model.Note)
	note.ID = params.NoteID
	core.Log.Info(params.NoteID)
	params.err = note.GetNote()
	params.result = gin.H{"title": note.Title, "preView": note.PreView, "content": note.Content}
}

type GetNoteList struct {
	baseParams `json:"-"`
	Page       int `form:"page" binding:"required"`
}

func (params *GetNoteList) Service() {
	notes := new(model.Notes)
	noteList, err := notes.GetNoteIDList(params.Page, core.Conf.GetNoteNumOfPage())
	if err != nil {
		params.err = err
		return
	}
	if len(noteList) == 0 {
		params.err = core.ErrEndOfNoteList
		return
	}
	params.result = gin.H{"noteList": noteList}
}
