package user

import (
	"github.com/ccchieh/ginHelper"
	"github.com/ccchieh/gospree/controller/restAPI/middleware"
	"github.com/ccchieh/gospree/model"
	"github.com/ccchieh/gospree/service"
	"github.com/ccchieh/gospree/util/ret"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary 通过ID获取用户信息
// @Tags 用户
// @Param id query int true "ID"
// @Success 200
// @Failure 400
// @Router /user [get]
// @version 1.0
func (h *Helper) GetUserInfoByIDHandler() (r *ginHelper.Router) {
	handler := func(c *gin.Context) {
		params := new(service.GetUserInfoByIDParams)
		if err := c.ShouldBind(params); err != nil {
			ret.Result(c, http.StatusBadRequest, nil, ret.ErrParameterMatch)
			return
		}
		value, exists := c.Get("user")
		if !exists {
			ret.Result(c, http.StatusBadRequest, nil, ret.ErrUserNotFound)
			return
		}
		user, ok := value.(*model.User)
		if !ok || user.ID != params.ID {
			ret.Result(c, http.StatusUnauthorized, nil, ret.ErrValidation)
			return
		}
		ret.Result(c, http.StatusOK, gin.H{"email": user.Email, "name": user.Name}, nil)
	}

	return &ginHelper.Router{
		Path:   "/",
		Method: "GET",
		Handlers: []gin.HandlerFunc{
			middleware.TokenMiddleware(),
			handler,
		}}
}

// @Tags 用户
// @Summary 新建用户
// @Param user body service.CreateUserParams true "用户"
// @Success 200
// @Failure 400
// @Router /user [post]
// @version 1.0
func (h *Helper) CreateUserHandler() (r *ginHelper.Router) {
	handler := func(c *gin.Context) {
		params := new(service.CreateUserParams)
		if err := c.ShouldBind(params); err != nil {
			ret.Result(c, http.StatusBadRequest, nil, ret.ErrParameterMatch)
			return
		}

		user, err := service.CreateUserService(params)
		if err != nil {
			ret.Result(c, http.StatusBadRequest, nil, err)
			return
		}
		ret.Result(c, http.StatusOK, gin.H{"email": user.Email, "name": user.Name}, nil)
	}

	return &ginHelper.Router{
		Path:   "/",
		Method: "POST",
		Handlers: []gin.HandlerFunc{
			handler,
		}}
}

// @Tags 用户
// @Summary 用户获取token
// @Param user body service.CreateTokenParams true "用户"
// @Success 200
// @Failure 400
// @Router /user/token [post]
// @version 1.0
func (h *Helper) CreateTokenHandler() (r *ginHelper.Router) {
	handler := func(c *gin.Context) {
		params := new(service.CreateTokenParams)
		if err := c.ShouldBind(params); err != nil {
			ret.Result(c, http.StatusBadRequest, nil, ret.ErrParameterMatch)
			return
		}

		token, id, err := service.CreateTokenService(params)
		if err != nil {
			ret.Result(c, http.StatusBadRequest, nil, err)
			return
		}
		ret.Result(c, http.StatusOK, gin.H{"id": id, "token": token}, nil)
	}

	return &ginHelper.Router{
		Path:   "/token",
		Method: "POST",
		Handlers: []gin.HandlerFunc{
			handler,
		}}
}
