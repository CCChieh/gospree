package service

import (
	"github.com/ccchieh/ginHelper"
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/util/ret"
	"github.com/gin-gonic/gin"
)

type Param struct {
	ginHelper.Param
}

func (param *Param) Bind(c *gin.Context, p ginHelper.Parameter) {
	if err := c.ShouldBind(p); err != nil {
		param.Err = core.ErrParameterMatch
	}
}

func (param *Param) Result(c *gin.Context) {
	ret.Result(c, param.Ret, param.Err)
}
