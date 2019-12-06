package user

import (
	"github.com/ccchieh/ginHelper"
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/dao"
	"github.com/ccchieh/gospree/util/ret"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 通过邮箱获取用户信息
// @Tags 用户
// @Param email query string true "邮箱"
// @Success 200 object dao.User "成功后返回值"
// @Failure 400 {string} string "fail" "失败后返回值"
// @Router /user [get]
// @version 1.0
func (h *Helper) GetUserInfoHandler() (r *ginHelper.Router) {
	return &ginHelper.Router{
		Path:   "/",
		Method: "GET",
		Handlers: []gin.HandlerFunc{
			getUserInfoByEmailHandler,
		}}
}
func getUserInfoByEmailHandler(c *gin.Context) {
	params := new(getUserInfoByEmailParams)
	if err := c.Bind(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, params)
}

// @Tags 用户
// @Summary 新建用户
// @Param user body user.createUserParams true "用户"
// @Success 200 object dao.User "成功后返回值"
// @Failure 400 {string} string "fail" "失败后返回值"
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
	user := new(dao.User)
	params := new(createUserParams)
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Email = params.Email
	user.Name = params.Name
	user.Password = params.Password
	core.Log.Info(user.Email, user.Name, user.Password)
	core.Log.Info(gin.H{"status": "you are logged in"})
	ret.Result(c, http.StatusOK, user, nil)
}
