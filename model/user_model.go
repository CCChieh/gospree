package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	Email      string `gorm:"type:varchar(64);unique_index;not null"`
	Name       string `gorm:"type:varchar(64);unique;not null"`
	Password   string `gorm:"type:varchar(64);not null"`
}

func (user *User) GetUserByEmail() error {
	return dao.First(user).Error
}

func (user *User) CreateUser() error {
	return dao.Create(user).Error
}

func (user *User) UserExist() bool {
	if err := dao.First(user).Error; err != nil {
		return false
	}
	return true
}
