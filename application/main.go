package application

import (
	"github.com/ccchieh/elog"
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/dao"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var App *application

func init() {
	App = new(application)
	elog.Logger.SetSave(true, "GoSpree.log")
	core.Log = elog.Logger
	core.Conf = new(config)
	core.Conf.SetDatabase(
		"gospree",
		"localhost",
		3306,
		"gospree",
		"wdsrshi19971025",
	)
	if err := dao.Connection(core.Conf.GetDatabaseUrl()); err != nil {
		core.Log.Err(err)
	}

}
