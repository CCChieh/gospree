package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var dao = new(db)

type db struct {
	*gorm.DB
}

func Connection(databaseUrl string) (err error) {
	if dao.DB, err = gorm.Open("mysql", databaseUrl); err != nil {
		return err
	}
	dao.LogMode(true)
	dao.AutoMigrate(&User{}, &Note{})
	return
}
