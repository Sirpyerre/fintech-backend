package repository

import "github.com/Sirpyerre/fintech-backend/internal/models"

type Transactioner interface {
	StoreTransaction(records []models.Transaction) error
}
