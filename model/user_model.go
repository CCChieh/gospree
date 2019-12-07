package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	Email      string `gorm:"type:varchar(64);unique_index"`
	Name       string `gorm:"type:varchar(64);unique;not null"`
	Password   string `gorm:"type:varchar(64);not null"`
}

func (user *User) GetUserByEmail(email string) error {
	return dao.Where(&User{Email: email}).First(&user).Error
}
