package repository

import (
	"context"

	"github.com/Sirpyerre/fintech-backend/internal/dbconnection"
	"github.com/Sirpyerre/fintech-backend/internal/models"
	"github.com/rs/zerolog"
)

type TransactionRepository struct {
	DBConnection *dbconnection.DBConnection
	logger       zerolog.Logger
}

func NewTransactionRepository(dbConnection *dbconnection.DBConnection, logger zerolog.Logger) *TransactionRepository {
	return &TransactionRepository{
		DBConnection: dbConnection,
		logger:       logger,
	}
}

func (r *TransactionRepository) StoreTransaction(ctx context.Context, records []models.Transaction) error {
	if len(records) == 0 {
		return nil
	}

	tx, err := r.DBConnection.Conn.Begin(ctx)
	if err != nil {
		r.logger.Error().Err(err).Msg("failed to begin transaction")
		return err
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				r.logger.Error().Err(rbErr).Msg("failed to rollback transaction")
			}
		}
	}()

	stmt := `INSERT INTO transactions (user_id, amount, datetime) VALUES ($1, $2, $3)`
	for _, t := range records {
		_, err = tx.Exec(ctx, stmt, t.UserID, t.Amount, t.OccurredAt)
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to insert transaction")
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				r.logger.Error().Err(rbErr).Msg("failed to rollback transaction after insert error")
			}
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		r.logger.Error().Err(err).Msg("failed to commit transaction")
		return err
	}
	return nil
}
