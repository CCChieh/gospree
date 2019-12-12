package ret

var (
	OK      = &retErr{Code: 0, Message: "OK"}
	Unknown = &retErr{Code: 99999, Message: "Unknown error"}
	// system errors

	ErrValidation     = &retErr{Code: 20001, Message: "Validation failed"}
	ErrParameterMatch = &retErr{Code: 20003, Message: "The parameters was not match"}

	// user errors
	ErrUserNotFound      = &retErr{Code: 20101, Message: "The user was not found"}
	ErrUserCreated       = &retErr{Code: 20102, Message: "The user was created"}
	ErrPasswordIncorrect = &retErr{Code: 20104, Message: "The password was incorrect."}
)

type retErr struct {
	Code    int
	Message string
}

func (err *retErr) Error() string {
	return err.Message
}
