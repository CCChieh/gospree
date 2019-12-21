package application

import (
	"fmt"
	"github.com/ccchieh/gospree/core"
	"github.com/ccchieh/gospree/util"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"os"
)

//配置文件
type config struct {
	Site struct {
		Name string
		Port int
		URL  string
		Note struct {
			NumOfPage int
		}
	}
	Database struct {
		Name     string
		Host     string
		Port     int
		User     string
		Password string
	}
	confFilePath string `json:"-"`
}

func (conf *config) SetSite(
	Name string,
	URL string,
	Port int) {
	conf.Site.Name = Name
	conf.Site.Port = Port
	conf.Site.URL = URL

	conf.Save()
}

func (conf *config) GetSiteName() (Name string) {
	return conf.Site.Name

}

func (conf *config) GetSiteURL() (URL string) {
	return conf.Site.URL
}

func (conf *config) GetSitePort() (Port int) {
	return conf.Site.Port
}

func (conf *config) SetNoteNumOfPage(NumOfPage int) {
	conf.Site.Note.NumOfPage = NumOfPage

	conf.Save()
}
func (conf *config) GetNoteNumOfPage() (NumOfPage int) {
	return conf.Site.Note.NumOfPage
}
func (conf *config) SetDatabase(
	Name string,
	Host string,
	Port int,
	User string,
	Password string,
) {
	conf.Database.Name = Name
	conf.Database.Host = Host
	conf.Database.Port = Port
	conf.Database.User = User
	conf.Database.Password = Password

	conf.Save()
}

func (conf *config) GetDatabaseUrl() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=True&loc=Local",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name)
}

func (conf *config) Save() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	//Convert to json.
	file, err := os.OpenFile(conf.confFilePath, os.O_CREATE, 0666)
	if err != nil {
		core.Log.Error(err)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	js, err := json.MarshalIndent(conf, "", " ")
	if err != nil {
		core.Log.Error(err)
		return
	}
	if _, err = file.Write(js); err != nil {
		core.Log.Error(err)
		return
	}
	return
}

func (conf *config) Init(confFilePath string) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	conf.confFilePath = confFilePath
	if util.FileIsExist(conf.confFilePath) {
		js, err := ioutil.ReadFile(conf.confFilePath)
		if err != nil {
			core.Log.Error(core.ErrLoadConfig)
		}
		err = json.Unmarshal(js, conf)
		if err != nil {
			core.Log.Error(core.ErrLoadConfig)
		}
		core.Panic(err)
	} else {
		conf.SetDatabase(
			"gospree",
			"lab.nimingshe.com",
			3306,
			"gospree",
			"wdsrshi19971025",
		)
		conf.SetSite(
			"Nimingshe",
			"localhost",
			8080)
		conf.SetNoteNumOfPage(10)
	}
}
