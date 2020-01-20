package model

import (
	"github.com/ccchieh/gospree/core"
	"github.com/jinzhu/gorm"
	"time"
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

type getNoteListResult struct {
	ID        uint
	Title     string
	PreView   string
	CreatedAt time.Time
	AuthorID  uint
	Author    string
}

//返回最新Note的list
func (note *Note) GetNoteList(page int, num int) ([]getNoteListResult, error) {

	var ret []getNoteListResult
	page -= 1
	err := dao.Model(&Note{}).Offset(page * num).Limit(num).Order("id desc").Scan(&ret).Error
	if len(ret) == 0 {
		err = core.ErrEndOfNoteList
		return ret, err
	}
	type Author struct {
		Name string
	}
	idNameMap := map[uint]string{}
	for i := range ret {
		if val, ok := idNameMap[ret[i].AuthorID]; ok {
			ret[i].Author = val
		} else {
			author := new(Author)
			dao.Model(&User{}).Select("name").Where("id = ?", ret[i].AuthorID).Scan(author)
			ret[i].Author = author.Name
			idNameMap[ret[i].AuthorID] = author.Name
		}

	}
	return ret, err
}
