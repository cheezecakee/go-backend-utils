package transaction

import (
	"context"
	"database/sql"

	"github.com/cheezecakee/go-backend-utils/pkg/logger"
)

//	type BaseRepository interface {
//		WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error
//
//		DB() *sql.DB
//	}
type BaseRepository struct {
	db *sql.DB
}

func NewBaseRepository(db *sql.DB) *BaseRepository {
	return &BaseRepository{db: db}
}

func (r *BaseRepository) WithTransaction(ctx context.Context, fn func(tx *sql.Tx) error) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		logger.Log.Error("Failed to begin transaction", map[string]any{
			"error": err.Error(),
		})
		return err
	}

	// Ensure the transaction is properly commited or rolled back
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			logger.Log.Critical("Panic in transaction", map[string]any{
				"panic": p,
			})
			panic(p)
		}
	}()

	// Execute the function with the transaction
	if err := fn(tx); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			logger.Log.Error("Failed to rollback transaction", map[string]any{
				"original_error": err.Error(),
				"rollback_error": rollbackErr.Error(),
			})
		}
		return err
	}

	return tx.Commit()
}

// func (r *baseRepository) DB() *sql.DB {
// 	return r.db
// }
