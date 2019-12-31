package user

import (
	"github.com/ccchieh/ginHelper"
	"github.com/ccchieh/gospree/controller/restAPI/middleware"
	"github.com/ccchieh/gospree/service"
	"github.com/gin-gonic/gin"
)

// @Summary 通过ID获取用户信息
// @Tags 用户
// @Param id query int true "ID"
// @Success 200
// @Failure 400
// @Router /user [get]
// @version 1.0
func (h *Helper) GetUserInfoHandler() (r *ginHelper.Router) {
	return &ginHelper.Router{
		Param:  new(service.GetUserInfo),
		Path:   "/",
		Method: "GET",
		Handlers: []gin.HandlerFunc{
			middleware.TokenMiddleware(),
		},
	}
}

// @Tags 用户
// @Summary 新建用户
// @Param user body service.CreateUser true "用户"
// @Success 200
// @Failure 400
// @Router /user [post]
// @version 1.0
func (h *Helper) CreateUserHandler() (r *ginHelper.Router) {
	return &ginHelper.Router{
		Param:  new(service.CreateUser),
		Path:   "/",
		Method: "POST",
	}
}

// @Tags 用户
// @Summary 用户获取token
// @Param user body service.CreateToken true "用户"
// @Success 200
// @Failure 400
// @Router /user/token [post]
// @version 1.0
func (h *Helper) CreateTokenHandler() (r *ginHelper.Router) {
	return &ginHelper.Router{
		Param:  new(service.CreateToken),
		Path:   "/token",
		Method: "POST",
	}
}
