package application

import (
	"github.com/ccchieh/gospree/controller/restAPI"
	"github.com/ccchieh/gospree/core"
)

type application struct {
}

func (app *application) Start() {
	core.Log.Info("Application has been starting")
	rest := new(restAPI.RestServer)
	rest.Start(core.Conf.GetSiteURL(), core.Conf.GetSitePort())
}

func (app *application) Close() {

}
