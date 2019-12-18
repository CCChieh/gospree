package core

type config interface {
	SetDatabase(
		Name string,
		Host string,
		Port int,
		User string,
		Password string,
	)
	GetDatabaseUrl() string
	Save()
	Init(confFilePath string)
}
