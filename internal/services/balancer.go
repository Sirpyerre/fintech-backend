package services

import (
	"context"
	"github.com/Sirpyerre/fintech-backend/internal/models"
	"time"
)

type Balancer interface {
	Balance(ctx context.Context, userID int64, from, to *time.Time) (models.BalanceResponse, error)
}
