package core

import (
	"fmt"
	"net/http"
)

var (
	OK            = &Err{statusCode: http.StatusOK, Code: 0, Message: "OK"}
	ErrUnknown    = &Err{statusCode: http.StatusInternalServerError, Code: 99999, Message: "ErrUnknown error"}
	ErrPermission = &Err{statusCode: http.StatusForbidden, Code: 10001, Message: "Need to check permissions"}
	ErrLoadConfig = &Err{statusCode: http.StatusInternalServerError, Code: 10002, Message: "Loading config file error"}
	// system errors

	ErrValidation     = &Err{statusCode: http.StatusUnauthorized, Code: 20001, Message: "Validation failed"}
	ErrParameterMatch = &Err{statusCode: http.StatusBadRequest, Code: 20003, Message: "Insufficient or non-compliant parameters provided"}

	// user errors
	ErrUserNotFound      = &Err{statusCode: http.StatusBadRequest, Code: 20101, Message: "The user was not found"}
	ErrUserCreated       = &Err{statusCode: http.StatusBadRequest, Code: 20102, Message: "The user was created"}
	ErrEndOfNoteList     = &Err{statusCode: http.StatusBadRequest, Code: 20102, Message: "The end of note list"}
	ErrPasswordIncorrect = &Err{statusCode: http.StatusForbidden, Code: 20104, Message: "The password was incorrect."}
)

type Err struct {
	statusCode int
	Code       int
	Message    string
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
