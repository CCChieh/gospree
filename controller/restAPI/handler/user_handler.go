package handler

import (
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 通过邮箱获取用户信息
// @Tags 用户
// @Param email query string true "用户的邮箱"
// @Success 200 object dao.User "成功后返回值"
// @Failure 400 {string} string "fail" "失败后返回值"
// @Router /user/info [get]
// @version 1.0
func GetUserInfoHandler(c *gin.Context) {
	email := c.Query("email")
	user := new(dao.User)
	if err := user.GetUserByEmail(email); err != nil {
		c.String(http.StatusBadRequest, "fail")
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Tags 用户
// @Summary 新建用户
// @Param user body handler.CreateUserParams true "用户"
// @Success 200 object dao.User "成功后返回值"
// @Failure 400 {string} string "fail" "失败后返回值"
// @Router /user [post]
// @version 1.0
func CreateUserHandler(c *gin.Context) {
	user := new(dao.User)
	createUserParam := new(CreateUserParams)
	if err := c.ShouldBindJSON(&createUserParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Email = createUserParam.Email
	user.Name = createUserParam.Name
	user.Password = createUserParam.Password
	core.Log.Info(user.Email, user.Name, user.Password)
	c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
