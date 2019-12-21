package service

import (
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/model"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(params *CreateUserParams) (user *model.User, err error) {
	user = &model.User{
		Email:    params.Email,
		Name:     params.Name,
		Password: params.Password,
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		core.Log.Err(err)
	}
	user.Password = string(hash)
	if err = user.CreateUser(); err != nil {
		err = core.ErrUserCreated
	}
	return user, err
}

func CreateTokenService(params *CreateTokenParams) (token string, err error) {
	user := &model.User{
		Email: params.Email,
	}
	if err = user.GetUser(); err != nil {
		err = core.ErrPasswordIncorrect
		return
	}
	if !user.Verification(params.Password) {
		err = core.ErrPasswordIncorrect
		return
	}
	token, err = user.CreateToken()
	return
}
