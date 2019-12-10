package service

import (
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/model"
	"github.com/ccchieh/gospree/util/ret"
	"golang.org/x/crypto/bcrypt"
)

func GetUserInfoByIDService(params *GetUserInfoByIDParams) (*model.User, error) {
	user := new(model.User)
	user.ID = params.ID
	err := user.GetUser()
	return user, err
}

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
		err = ret.ErrUserCreated
	}
	return user, err
}

func CreateTokenService(params *CreateTokenParams) (token string, id uint, err error) {
	user := &model.User{
		Email: params.Email,
	}
	if err = user.GetUser(); err != nil {
		err = ret.ErrPasswordIncorrect
		return
	}
	id = user.ID
	if !user.Verification(params.Password) {
		err = ret.ErrPasswordIncorrect
	}
	token, err = user.CreateToken()
	return
}
