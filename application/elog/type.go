package elog

const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
)

type Err interface {
	error
	GetLevel() (level int)
	Local() (file string, line int)
}

type errBase struct {
	Msg   string `json:"msg"`
	Level int    `json:"-"`
	File  string `json:"-"`
	Line  int    `json:"-"`
}

func (err *errBase) Error() string {
	return err.Msg
}

func (err *errBase) GetLevel() int {
	return err.Level
}
func (err *errBase) Local() (file string, line int) {
	return err.File, err.Line
}

type errPathNotExist struct {
	*errBase
	Path string `json:"-"`
}

type errPathNotDir struct {
	*errBase
	Path string `json:"-"`
}

type errPathNotFile struct {
	*errBase
	Path string `json:"-"`
}

type errPathNameIsExist struct {
	*errBase
	Path string `json:"-"`
	Name string `json:"-"`
}

type errUserNameIsExist struct {
	*errBase
	Name string `json:"-"`
}

type errUserEmailIsExist struct {
	*errBase
	Email string `json:"-"`
}

type errLoginFail struct {
	*errBase
	Email string `json:"-"`
}

type errEmailFormat struct {
	*errBase
	Email string `json:"-"`
}

type errUserNameFormat struct {
	*errBase
	UserName string `json:"-"`
}

type errPasswordFormat struct {
	*errBase
	Password string `json:"-"`
}

type errUnauthorized struct {
	*errBase
	Path string `json:"-"`
}
