package service

import (
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/model"
)

func GetNoteService(params *GetNoteParams) (note *model.Note, err error) {
	note = new(model.Note)
	note.ID = params.NoteID
	err = note.GetNote()
	return
}

func CreateNoteService(params *CreateNoteParams) (note *model.Note, err error) {
	note = &model.Note{
		Title:    params.Title,
		Content:  params.Content,
		AuthorID: params.ID,
	}
	err = note.CreateNote()
	return
}

func GetNoteIDListService(params *GetNoteIDListParams) (noteIDs []uint, err error) {
	notes := new(model.Notes)
	noteIDs, err = notes.GetNoteIDList(params.Page, core.Conf.GetNoteNumOfPage())
	if err != nil {
		return
	}
	if len(noteIDs) == 0 {
		err = core.ErrEndOfNoteList
	}
	return
}
