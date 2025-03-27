package errors

import (
	"fmt"
	"runtime"
)

type CustomError struct {
	Code    string
	Message string
	Err     error
	Stack   string
}

func New(code, message string) *CustomError {
	stack := captureStackTrace()
	return &CustomError{
		Code:    code,
		Message: message,
		Stack:   stack,
	}
}

func Wrap(err error, code, message string) *CustomError {
	stack := captureStackTrace()
	return &CustomError{
		Code:    code,
		Message: message,
		Err:     err,
		Stack:   stack,
	}
}

func (e *CustomError) UnWrap() error {
	return e.Err
}

func captureStackTrace() string {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, false)
	return string(buf[:n])
}

func (e *CustomError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s [stack: %s]\nUnWrapped: %v", e.Code, e.Message, e.Stack, e.Err)
	}
	return fmt.Sprintf("%s: %s [stack: %s]", e.Code, e.Message, e.Stack)
}
