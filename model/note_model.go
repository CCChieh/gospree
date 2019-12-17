package model

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model `json:"-"`
	Content    string
	AuthorID   uint
}

func (note *Note) GetNote() error {
	return dao.Where(note).First(note).Error
}

func (note *Note) CreateNote() error {
	return dao.Create(note).Error
}
