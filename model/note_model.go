package model

import (
	"github.com/jinzhu/gorm"
)

type Note struct {
	gorm.Model `json:"-"`
	Title      string
	Content    string `gorm:"type:text"`
	PreView    string `gorm:"type:varchar(255)"`
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
	var count int
	page -= 1
	err := dao.Model(&Note{}).Count(&count).Offset(page*num).Limit(num).Order("id desc").Pluck("id", &notesID).Error
	return notesID, err
}
