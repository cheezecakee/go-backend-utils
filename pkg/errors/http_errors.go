package errors

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/cheezecakee/go-backend-utils/pkg/logger"
)

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	logger.Log.Error(trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func ClientError(w http.ResponseWriter, status int) {
	logger.Log.Warn(fmt.Sprintf("Client error: %d - %s", status, http.StatusText(status)))
	http.Error(w, http.StatusText(status), status)
}

func NotFound(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}
