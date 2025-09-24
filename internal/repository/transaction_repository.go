package repository

import (
	"context"
	"github.com/Sirpyerre/fintech-backend/internal/dbconnection"
	"github.com/Sirpyerre/fintech-backend/internal/models"
	"github.com/rs/zerolog"
	"time"
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

	// Defer rollback only if commit hasn't happened
	committed := false
	defer func() {
		if !committed {
			if rbErr := tx.Rollback(ctx); rbErr != nil {
				r.logger.Error().Err(rbErr).Msg("failed to rollback transaction")
			}
		}
	}()

	stmt := `INSERT INTO transactions (user_id, amount, datetime) VALUES ($1, $2, $3)`
	for _, t := range records {
		if _, execErr := tx.Exec(ctx, stmt, t.UserID, t.Amount, t.OccurredAt); execErr != nil {
			r.logger.Error().Err(execErr).Msg("failed to insert transaction")
			return execErr
		}
	}

	if commitErr := tx.Commit(ctx); commitErr != nil {
		r.logger.Error().Err(commitErr).Msg("failed to commit transaction")
		return commitErr
	}
	committed = true
	return nil
}

func (r *TransactionRepository) FindUserExists(ctx context.Context, userID int64) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)`
	err := r.DBConnection.Conn.QueryRow(ctx, query, userID).Scan(&exists)
	if err != nil {
		r.logger.Error().Err(err).Msg("the user does not exist")
		return false, err
	}

	return exists, nil
}

func (r *TransactionRepository) BalanceSummary(ctx context.Context, userID int64, from, to *time.Time) (float64, float64, float64, error) {
	var totalCredits, totalDebits float64
	query := `SELECT 
				COALESCE(SUM(CASE WHEN amount > 0 THEN amount ELSE 0 END), 0) AS total_credits,
				COALESCE(SUM(CASE WHEN amount < 0 THEN amount ELSE 0 END), 0) AS total_debits
			  FROM transactions 
			  WHERE user_id = $1`

	if from != nil && to != nil {
		query += ` AND datetime BETWEEN $2 AND $3`
		err := r.DBConnection.Conn.QueryRow(ctx, query, userID, *from, *to).Scan(&totalCredits, &totalDebits)
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to get balance summary with date range")
			return 0, 0, 0, err
		}
	} else {
		err := r.DBConnection.Conn.QueryRow(ctx, query, userID).Scan(&totalCredits, &totalDebits)
		if err != nil {
			r.logger.Error().Err(err).Msg("failed to get balance summary without date range")
			return 0, 0, 0, err
		}
	}

	balance := totalCredits + totalDebits
	return totalCredits, totalDebits, balance, nil
}
