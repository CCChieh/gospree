package service

import (
	"github.com/ccchieh/gospree/model"
	"github.com/ccchieh/gospree/util/ret"
)

func GetUserInfoByEmailService(params *GetUserInfoByEmailParams) (*model.User, error) {
	user := &model.User{
		Email: params.Email,
	}
	err := user.GetUserByEmail()
	return user, err
}

func CreateUserService(params *CreateUserParams) (user *model.User, err error) {
	user = &model.User{
		Email:    params.Email,
		Name:     params.Name,
		Password: params.Password,
	}
	if err = user.CreateUser(); err != nil {
		err = ret.ErrUserCreated
	}
	return user, err
}
