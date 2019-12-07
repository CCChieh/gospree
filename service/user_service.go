package service

import "github.com/ccchieh/gospree/model"

func GetUserInfoByEmailService(email string) (*model.User, error) {
	user := new(model.User)
	err := user.GetUserByEmail(email)
	return user, err
}
