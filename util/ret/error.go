package ret

var (
	OK      = &retErr{Code: 0, Message: "OK"}
	Unknown = &retErr{Code: 99999, Message: "Unknown error"}
	// system errors
	ErrInternalServer = &retErr{Code: 10001, Message: "Internal server error"}
	ErrBind           = &retErr{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &retErr{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &retErr{Code: 20002, Message: "Database error."}
	ErrToken      = &retErr{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrUserNotFound      = &retErr{Code: 20101, Message: "The user was not found."}
	ErrUserCreated       = &retErr{Code: 20102, Message: "The user was created."}
	ErrPasswordIncorrect = &retErr{Code: 20104, Message: "The password was incorrect."}
)

type retErr struct {
	Code    int
	Message string
}

func (err *retErr) Error() string {
	return err.Message
}
