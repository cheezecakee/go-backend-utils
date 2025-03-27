package service

import (
	"context"
	"database/sql"

	m "github/cheezecakee/fitrkr/internal/user/models"
	u "github/cheezecakee/go-backend-utils/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

func (s *DBUserService) Register(ctx context.Context, user *m.User) (*m.User, error) {
	// Check for existing Email
	existingUser, err := s.repo.GetByEmail(ctx, user.Email)
	if err != nil && err != sql.ErrNoRows {
		u.Log.ErrorLog.Printf("Database error while checking email %s: %v", user.Email, err)
		return nil, u.errors.ErrInternalServer(err)
	}
	if existingUser != nil {
		u.Log.ErrorLog.Printf("Email already exists: %s", user.Email)
		return nil, u.errors.New(u.errors.ErrEmailExists, "Email already exists")
	}

	// Check for existing Username
	existingUser, err = s.repo.GetByUsername(ctx, user.Username)
	if err != nil && err != sql.ErrNoRows {
		u.Log.ErrorLog.Printf("Database error while checking username %s: %v", user.Username, err)
		return nil, u.Errors.ErrInternalServer(err)
	}
	if existingUser != nil {
		u.Log.ErrorLog.Printf("Username already taken: %s", user.Username)
		return nil, u.errors.New(u.errors.ErrUsernameTaken, "Username alrady taken")
	}

	hashedPassword, err := u.helper.PasswordHash(user.PasswordHash)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(hashedPassword)

	return s.repo.Create(ctx, user)
}
