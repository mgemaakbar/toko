package customerror

import "net/http"

type InternalError struct {
	msg string
}

func NewInternalError(msg string) *InternalError {
	return &InternalError{msg: msg}
}

func (i *InternalError) Error() string {
	return i.msg
}

type UserError struct {
	msg string
}

func NewUserError(msg string) *UserError {
	return &UserError{msg: msg}
}

func (i *UserError) Error() string {
	return i.msg
}

func GetHTTPResponseCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	switch err.(type) {
	case *InternalError:
		return http.StatusInternalServerError
	case *UserError:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}
