package service

type CreateNoteParams struct {
	Content  string `json:"content" binding:"required"`
	AuthorID uint   `json:"authorID" binding:"required"`
}

type GetNoteParams struct {
	NoteID uint `form:"noteID" binding:"required"`
}
