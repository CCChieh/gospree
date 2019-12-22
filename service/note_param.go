package service

type CreateNoteParams struct {
	Title   string `json:"title" binding:"required"`
	PreView string `json:"preView" binding:"require"`
	Content string `json:"content" binding:"required"`
	ID      uint   `json:"-"`
}

type GetNoteParams struct {
	NoteID uint `form:"noteID" binding:"required"`
}

type GetNoteIDListParams struct {
	Page int `form:"page" binding:"required"`
}
