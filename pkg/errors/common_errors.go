package errors

// Some common codes that can be used

const (
	// User-related errors
	ErrStatusNotFound     = "NOT_FOUND"
	ErrUserNotFound       = "USER_NOT_FOUND"
	ErrInvalidCredentials = "INVALID_CREDENTIALS"
	ErrInvalidUsername    = "INVALID_USERNAME_FORMAT"
	ErrWeakPassword       = "WEAK_PASSWORD"
	ErrEmailExists        = "EMAIL_ALREADY_EXISTS"
	ErrInvalidEmail       = "INVALID_EMAIL_FORMAT"

	// Backend errors
	ErrDBError         = "DB_ERROR"
	ErrValidationError = "VALIDATION_ERROR"

	// General errors
	ErrBadRequest          = "BAD_REQUEST"
	ErrConflict            = "CONFLICT_DETECTED"
	ErrOrderStatusNotFound = "ORDER_STATUS_NOT_FOUND"

	// Authentication & Authorization errors
	ErrUnauthorized   = "UNAUTHORIZED"
	ErrForbidden      = "FORBIDDEN"
	ErrInternalServer = "INTERNAL_SERVER_ERROR"
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
