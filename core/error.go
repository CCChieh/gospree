package core

import "fmt"

var (
	OK            = &Err{Code: 0, Message: "OK"}
	ErrUnknown    = &Err{Code: 99999, Message: "ErrUnknown error"}
	ErrPermission = &Err{Code: 10001, Message: "Need to check permissions"}
	ErrLoadConfig = &Err{Code: 10002, Message: "Loading config file error"}
	// system errors

	ErrValidation     = &Err{Code: 20001, Message: "Validation failed"}
	ErrParameterMatch = &Err{Code: 20003, Message: "Insufficient or non-compliant parameters provided"}

	// user errors
	ErrUserNotFound      = &Err{Code: 20101, Message: "The user was not found"}
	ErrUserCreated       = &Err{Code: 20102, Message: "The user was created"}
	ErrEndOfNoteList     = &Err{Code: 20102, Message: "The end of note list"}
	ErrPasswordIncorrect = &Err{Code: 20104, Message: "The password was incorrect."}
)

type Err struct {
	Code    int
	Message string
}

func (err *Err) Error() string {
	return err.Message
}

func Panic(err error) {
	if err == nil {
		return
	}
	if Log == nil {
		fmt.Println(err)
	} else {
		Log.Panic(err)
	}
	panic(err)
}
