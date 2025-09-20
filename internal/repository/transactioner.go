package repository

import (
	"context"
	"github.com/Sirpyerre/fintech-backend/internal/models"
)

type Transactioner interface {
	StoreTransaction(ctx context.Context, records []models.Transaction) error
}
