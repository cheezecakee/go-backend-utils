package logger

import (
	"context"

	"github.com/google/uuid"
)

// Context key to prevent key conflicts
type ContextKey string

// Context for tracing and logging
var (
	RequestIDKey = ContextKey("request_id")
	UserIDKey    = ContextKey("user_id")
)

// Adds a request key to the context
func WithRequestID(ctx context.Context) context.Context {
	return context.WithValue(ctx, RequestIDKey, uuid.New().String())
}

func GetRequestID(ctx context.Context) string {
	if requestID, ok := ctx.Value(RequestIDKey).(string); ok {
		return requestID
	}
	return "unknown"
}
