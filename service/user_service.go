package service

import (
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type CreateUser struct {
	Param
	Email    string `json:"email" binding:"required,email"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (param *CreateUser) Service() {
	user := &model.User{
		Email:    param.Email,
		Name:     param.Name,
		Password: param.Password,
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		param.Err = core.ErrUserCreated
		return
	}
	user.Password = string(hash)
	if err = user.CreateUser(); err != nil {
		param.Err = core.ErrUserCreated
		return
	}
	param.Ret = gin.H{"email": user.Email, "name": user.Name}
}

type GetUserInfo struct {
	Param
}

func (param *GetUserInfo) BeforeBind(c *gin.Context) {
	value, exists := c.Get("user")
	if !exists {
		param.Err = core.ErrPermission
		return
	}
	user := value.(*model.User)
	param.Ret = gin.H{"email": user.Email, "name": user.Name}
}

type CreateToken struct {
	Param
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (param *CreateToken) Service() {
	user := &model.User{
		Email: param.Email,
	}
	if err := user.GetUser(); err != nil {
		param.Err = core.ErrPasswordIncorrect
		return
	}
	if !user.Verification(param.Password) {
		param.Err = core.ErrPasswordIncorrect
		return
	}
	if token, err := user.CreateToken(); err != nil {
		param.Err = err
	} else {
		param.Ret = gin.H{"token": token}
	}
}
