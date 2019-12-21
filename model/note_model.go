package model

import (
	"github.com/ccchieh/gospree/core"
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model `json:"-"`
	Title      string
	Content    string `gorm:"type:text"`
	AuthorID   uint
}

func (note *Note) GetNote() error {
	return dao.Where(note).First(note).Error
}

func (note *Note) CreateNote() error {
	return dao.Create(note).Error
}

type Notes []Note

//返回最新Note的ID
func (notes *Notes) GetNoteIDList(page int, num int) ([]uint, error) {
	var notesID []uint
	page -= 1
	core.Log.Info(page, num)
	err := dao.Offset(page*num).Limit(num).Model(&Note{}).Order("id desc").Pluck("id", &notesID).Error
	return notesID, err
}
