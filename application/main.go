package application

import (
	"github.com/ccchieh/elog"
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/model"
)

var App *application

func init() {
	App = new(application)
	elog.Logger.SetSave(true, "GoSpree.log")
	core.Log = elog.Logger
	core.Conf = new(config)
	core.Conf.Init("config.json")
	if err := model.Connection(core.Conf.GetDatabaseUrl()); err != nil {
		core.Log.Err(err)
	}

}
