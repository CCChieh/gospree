package service

import (
	"github.com/ccchieh/gospree/controller/restAPI"
	"github.com/ccchieh/gospree/core"
	"reflect"
)

type application struct {
}

func (app *application) Start() {
	core.Log.Info("Application has been starting")
	test()
	rest := new(restAPI.RestServer)
	rest.Start(8080)
}

func (app *application) Close() {

}

type myType struct {
	name string
	id   int
}

func (t *myType) Hello() {
	core.Log.Info("Hello p")
}

func (t *myType) String() string {
	core.Log.Info("Hello String")
	return "Hello"
}

func (t *myType) GetName(in string) {
	core.Log.Info(t.name, in)
}

func (t *myType) GetId(in string) {
	core.Log.Info(t.id, in)
}

func Hello(name string) string {
	core.Log.Info("Hello ", name)
	return name
}
func test() {
	myType := &myType{"Chieh", 12}
	mtV := reflect.ValueOf(myType)
	core.Log.Info(mtV)
	mtV.MethodByName("Hello").Call(nil)
	core.Log.Info(mtV.MethodByName("String").Call(nil)[0])
}
