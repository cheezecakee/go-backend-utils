package helper

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/cheezecakee/go-backend-utils/pkg/errors"
	"github.com/cheezecakee/go-backend-utils/pkg/logger"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ValidatePasswordHash(hash, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return err
	}
	return nil
}

func GetBearerToken(headers http.Header) (string, error) {
	token := headers.Get("Authorization")
	if token == "" {
		logger.Log.Warn("missing authorization header")
		return "", errors.New(errors.ErrUnauthorized, "missing authorization header")

	}
	token = token[len("Bearer "):]
	return token, nil
}

func MakeRefreshToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		logger.Log.Error("error generating a new token")
		return "", errors.New(errors.ErrInternalServer, "error generating a new token")
	}

	encodedToken := hex.EncodeToString(token)
	return encodedToken, nil
}

func Clamp(value, min, max int) int {
	switch {
	case value < min:
		return min
	case value > max:
		return max
	default:
		return value
	}
}
