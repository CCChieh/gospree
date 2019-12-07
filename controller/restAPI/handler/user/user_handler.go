package user

import (
	"github.com/ccchieh/ginHelper"
	"github.com/ccchieh/gospree/service"
	"github.com/ccchieh/gospree/util/ret"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 通过邮箱获取用户信息
// @Tags 用户
// @Param email query string true "邮箱"
// @Success 200
// @Failure 400
// @Router /user [get]
// @version 1.0
func (h *Helper) GetUserInfoByEmailHandler() (r *ginHelper.Router) {
	return &ginHelper.Router{
		Path:   "/",
		Method: "GET",
		Handlers: []gin.HandlerFunc{
			getUserInfoByEmailHandler,
		}}
}
func getUserInfoByEmailHandler(c *gin.Context) {
	params := new(service.GetUserInfoByEmailParams)
	if err := c.Bind(params); err != nil {
		ret.Result(c, http.StatusBadRequest, nil, err)
		return
	}
	user, err := service.GetUserInfoByEmailService(params)
	if err != nil {
		ret.Result(c, http.StatusBadRequest, nil, ret.ErrUserNotFound)
		return
	}
	ret.Result(c, http.StatusOK, gin.H{"email": user.Email, "name": user.Name}, nil)
}

// @Tags 用户
// @Summary 新建用户
// @Param user body service.CreateUserParams true "用户"
// @Success 200
// @Failure 400
// @Router /user [post]
// @version 1.0
func (h *Helper) CreateUserHandler() (r *ginHelper.Router) {
	return &ginHelper.Router{
		Path:   "/",
		Method: "POST",
		Handlers: []gin.HandlerFunc{
			createUserHandler,
		}}
}

func createUserHandler(c *gin.Context) {
	params := new(service.CreateUserParams)
	if err := c.Bind(&params); err != nil {
		ret.Result(c, http.StatusBadRequest, nil, err)
		return
	}

	user, err := service.CreateUserService(params)
	if err != nil {
		ret.Result(c, http.StatusBadRequest, nil, err)
		return
	}
	ret.Result(c, http.StatusOK, gin.H{"email": user.Email, "name": user.Name}, nil)
}
