package repository

import (
	"context"
	"github.com/Sirpyerre/fintech-backend/internal/models"
	"time"
)

type Transactioner interface {
	StoreTransaction(ctx context.Context, records []models.Transaction) error
	FindUserExists(ctx context.Context, userID int64) (bool, error)
	BalanceSummary(ctx context.Context, userID int64, from, to *time.Time) (float64, float64, float64, error)
}
