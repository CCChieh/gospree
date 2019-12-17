package service

import (
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
		Content:  params.Content,
		AuthorID: params.AuthorID,
	}
	err = note.CreateNote()
	return
}
