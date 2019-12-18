package service

type CreateNoteParams struct {
	Content string `json:"content" binding:"required"`
	ID      uint
}

type GetNoteParams struct {
	NoteID uint `form:"noteID" binding:"required"`
}
