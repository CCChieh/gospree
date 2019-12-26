package service

import (
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/model"
	"github.com/ccchieh/gospree/util/ret"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Params interface {
	Err() error
	Result() interface{}
	Service()
}

type baseParams struct {
	err    error
	result interface{}
}

func (base *baseParams) Err() error {
	return base.err
}

func (base *baseParams) Result() interface{} {
	return base.result
}
func (base *baseParams) Service() {
	//默认不做任何操作
}
func Process(c *gin.Context, params Params) {
	if err := c.ShouldBind(params); err != nil {
		ret.Result(c, http.StatusBadRequest, nil, core.ErrParameterMatch)
		return
	}
	params.Service()
	ret.Result(c, http.StatusOK, params.Result(), params.Err())
}

func GetID(c *gin.Context) uint {
	value, exists := c.Get("user")
	if !exists {
		ret.Result(c, http.StatusBadRequest, nil, core.ErrPermission)
		return 0
	}
	user := value.(*model.User)
	return user.ID
}
