package service

import "errors"

var (
	ErrLoginIsNotUnique = errors.New("login is not unique")
	ErrTokenExpired     = errors.New("token is expired")
)

type ValidationError struct {
	msg       string
	errorsMap map[string]string
}

func NewValidationError(message string, errorsMap map[string]string) error {
	return ValidationError{
		msg:       message,
		errorsMap: errorsMap,
	}
}

func (v ValidationError) Error() string {
	return v.msg
}

func (v ValidationError) ErrorsMap() map[string]string {
	return v.errorsMap
}
