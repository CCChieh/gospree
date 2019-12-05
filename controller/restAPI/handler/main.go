package handler

import (
	"github.com/ccchieh/ginHelper"
	"github.com/gin-gonic/gin"
	"reflect"
)

type helper struct {
}

func Build(r gin.IRoutes) {
	h := new(helper)
	valueOfh := reflect.ValueOf(h)
	numMethod := valueOfh.NumMethod()
	for i := 0; i < numMethod; i++ {
		rt := valueOfh.Method(i).Call(nil)[0].Interface().(*ginHelper.Router)
		rt.AddHandler(r)
	}
}
