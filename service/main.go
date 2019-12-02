package service

import (
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/service/elog"
)

var App *application

func init() {
	App = new(application)
	elog.Logger.SetSave(true, "GoSpree.log")
	core.Log = elog.Logger
}
