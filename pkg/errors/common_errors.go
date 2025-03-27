package errors

// Some common codes that can be used

const (
	ErrNotFound        = "NOT_FOUND"
	ErrUserNotFound    = "USER_NOT_FOUND"
	ErrStatusNotFound  = "STATUS_NOT_FOUND"
	ErrDBError         = "DB_ERROR"
	ErrValidationError = "VALIDATION_ERROR"
	ErrUnauthorized    = "UNAUTHORIZED"
	ErrForbidden       = "FORBIDDEN"
	ErrInternalServer  = "INTERNAL_SERVER_ERROR"
)

func UserNotFound(message string) *CustomError {
	return New(ErrUserNotFound, message)
}

func StatusNotFound(message string) *CustomError {
	return New(ErrStatusNotFound, message)
}

func DBError(err error, message string) *CustomError {
	return Wrap(err, ErrDBError, message)
}

func ValidationError(message string) *CustomError {
	return New(ErrValidationError, message)
}

func UnauthorizedError(message string) *CustomError {
	return New(ErrUnauthorized, message)
}

func Forbidden(message string) *CustomError {
	return New(ErrForbidden, message)
}

func InternalServerError(message string) *CustomError {
	return New(ErrInternalServer, message)
}
