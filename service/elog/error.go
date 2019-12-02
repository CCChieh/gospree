package elog

import (
	"fmt"
	"runtime"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func newErr() *errBase {
	err := new(errBase)
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	err.File, err.Line = file, line
	return err
}

func WrapErr(err error) error {
	base := newErr()
	base.Msg = err.Error()
	base.Level = ErrorLevel
	return base
}

func NewErrPathNotExist(path string) error {
	err := errPathNotExist{newErr(), path}
	err.Level = WarnLevel
	err.Msg = fmt.Sprintf("\"%s\" does not exist", err.Path)
	return err
}

func NewErrPathNotDir(path string) error {
	err := errPathNotDir{newErr(), path}
	err.Level = WarnLevel
	err.Msg = fmt.Sprintf("\"%s\" is not a directory", err.Path)
	return err
}

func NewErrPathNotFile(path string) error {
	err := errPathNotFile{newErr(), path}
	err.Level = WarnLevel
	err.Msg = fmt.Sprintf("\"%s\" is not a file", err.Path)
	return err
}

func NewErrPathNameIsExist(path string, name string) error {
	err := errPathNameIsExist{newErr(), path, name}
	err.Level = WarnLevel
	err.Msg = fmt.Sprintf("\"%s\" already exists", err.Path)
	return err
}
func NewErrUserNameIsExist(name string) error {
	err := errUserNameIsExist{newErr(), name}
	err.Level = WarnLevel
	err.Msg = fmt.Sprintf("Name:\"%s\" already exists", err.Name)
	return err
}
func NewErrUserEmailIsExist(email string) error {
	err := errUserEmailIsExist{newErr(), email}
	err.Level = WarnLevel
	err.Msg = fmt.Sprintf("Email:\"%s\" already exists", err.Email)
	return err
}

func NewErrLoginFail(email string) error {
	err := errLoginFail{newErr(), email}
	err.Level = WarnLevel
	err.Msg = fmt.Sprintf("Email:\"%s\" login fail", err.Email)
	return err
}

func NewErrEmailFormat(email string) error {
	err := errEmailFormat{newErr(), email}
	err.Level = WarnLevel
	err.Msg = fmt.Sprintf("Email:\"%s\" incorrect format", err.Email)
	return err
}
func NewErrUserNameFormat(userName string) error {
	err := errUserNameFormat{newErr(), userName}
	err.Level = WarnLevel
	err.Msg = fmt.Sprintf("User Name:\"%s\" incorrect format", err.UserName)
	return err
}

func NewErrPasswordFormat(password string) error {
	err := errPasswordFormat{newErr(), password}
	err.Level = WarnLevel
	err.Msg = fmt.Sprintf("Password:\"%s\" incorrect format", err.Password)
	return err
}

func NewErrUnauthorized(path string) error {
	err := errUnauthorized{newErr(), path}
	err.Level = WarnLevel
	err.Msg = fmt.Sprintf("Path:\"%s\" unauthorized", err.Path)
	return err
}
