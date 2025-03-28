package util

import (
	"github.com/cheezecakee/go-backend-utils/pkg/errors"
	"github.com/cheezecakee/go-backend-utils/pkg/helper"
	"github.com/cheezecakee/go-backend-utils/pkg/logger"
	"github.com/cheezecakee/go-backend-utils/pkg/middleware"
	"github.com/cheezecakee/go-backend-utils/pkg/transaction"
)

var (
	// Logger
	Log = logger.Log
	// Context
	WithRequestID = logger.WithRequestID
	GetRequestID  = logger.GetRequestID

	// Error handling
	NewError  = errors.New
	WrapError = errors.Wrap
	// Common error types
	UserNotFound        = errors.UserNotFound
	StatusNotFound      = errors.StatusNotFound
	DBError             = errors.DBError
	ValidationError     = errors.ValidationError
	UnauthorizedError   = errors.UnauthorizedError
	Forbidden           = errors.Forbidden
	InternalServerError = errors.InternalServerError
	// Common http error type
	ServerError = errors.ServerError
	ClientError = errors.ClientError
	NotFound    = errors.NotFound

	// Helper functions
	HashPassword         = helper.HashPassword
	ValidatePasswordHash = helper.ValidatePasswordHash
	GetBearerToken       = helper.GetBearerToken
	MakeRefreshToken     = helper.MakeRefreshToken
	Clamp                = helper.Clamp

	// Middlerware
	LoggingMiddleware = middleware.LogginMiddleware

	// Transaction
	NewBaseRepository = transaction.NewBaseRepository
	BaseRepository    = transaction.BaseRepository{}
)

const (
	ErrStatusNotFound     = errors.ErrStatusNotFound
	ErrUserNotFound       = errors.ErrUserNotFound
	ErrInvalidCredentials = errors.ErrInvalidCredentials
	ErrInvalidUsername    = errors.ErrInvalidUsername
	ErrWeakPassword       = errors.ErrWeakPassword
	ErrEmailExists        = errors.ErrEmailExists
	ErrInvalidEmail       = errors.ErrInvalidEmail
	ErrDBError            = errors.ErrDBError
	ErrValidationError    = errors.ErrValidationError
	ErrBadRequest         = errors.ErrBadRequest
	ErrConflict           = errors.ErrConflict
	ErrUnauthorized       = errors.ErrUnauthorized
	ErrForbidden          = errors.ErrForbidden
	ErrInternalServer     = errors.ErrInternalServer
)
