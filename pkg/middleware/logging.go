package middleware

import (
	"net/http"
	"time"

	"github.com/cheezecakee/go-backend-utils/pkg/logger"
)

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		ctx := logger.WithRequestID(r.Context())
		r = r.WithContext(ctx)

		wrapper := &responseWriterWrapper{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapper, r)

		logger.Log.Info("Request processed", map[string]any{
			"method":     r.Method,
			"path":       r.URL.Path,
			"status":     wrapper.statusCode,
			"duration":   time.Since(start).String(),
			"request_id": logger.GetRequestID(ctx),
			"remote_up":  r.RemoteAddr,
		})
	})
}

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
