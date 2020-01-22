package core

import (
	"fmt"
	"net/http"
)

var (
	OK            = &Err{HttpStatusCode: http.StatusOK, Code: 0, Message: "OK"}
	ErrUnknown    = &Err{HttpStatusCode: http.StatusInternalServerError, Code: 99999, Message: "ErrUnknown error"}
	ErrPermission = &Err{HttpStatusCode: http.StatusForbidden, Code: 10001, Message: "Need to check permissions"}
	ErrLoadConfig = &Err{HttpStatusCode: http.StatusInternalServerError, Code: 10002, Message: "Loading config file error"}
	// system errors

	ErrValidation     = &Err{HttpStatusCode: http.StatusUnauthorized, Code: 20001, Message: "Validation failed"}
	ErrParameterMatch = &Err{HttpStatusCode: http.StatusBadRequest, Code: 20003, Message: "Insufficient or non-compliant parameters provided"}

	// user errors
	ErrUserNotFound      = &Err{HttpStatusCode: http.StatusNotFound, Code: 20101, Message: "The user was not found"}
	ErrUserCreated       = &Err{HttpStatusCode: http.StatusBadRequest, Code: 20102, Message: "The user was created"}
	ErrEndOfNoteList     = &Err{HttpStatusCode: http.StatusNotFound, Code: 20103, Message: "The end of note list"}
	ErrPasswordIncorrect = &Err{HttpStatusCode: http.StatusForbidden, Code: 20104, Message: "The password was incorrect."}

	//note errors
	ErrNoteNotFound = &Err{HttpStatusCode: http.StatusNotFound, Code: 20201, Message: "The note was not found"}
)

type Err struct {
	HttpStatusCode int
	Code           int
	Message        string
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
