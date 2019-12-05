package application

import "fmt"

//配置文件
type config struct {
	Site struct {
		Name        string
		Port        int
		LogSaveDays int
	}
	Database struct {
		Name     string
		Host     string
		Port     int
		User     string
		Password string
	}
	status struct {
		installed bool
		loaded    bool
	}
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
}

func (conf *config) GetDatabaseUrl() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?parseTime=True&loc=Local",
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Host,
		conf.Database.Port,
		conf.Database.Name)
}

func (conf *config) IsInstall() bool {
	return false
}
