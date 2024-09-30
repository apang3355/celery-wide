package errs

import "fmt"

type Err string

func (e Err) Error() string {
	return string(e)
}

func NewError(msg string, fields ...any) error {
	return Err(fmt.Sprintf(msg, fields...))
}

func NewErrorMessage(msg string, fields ...any) string {
	return NewError(msg, fields...).Error()
}
