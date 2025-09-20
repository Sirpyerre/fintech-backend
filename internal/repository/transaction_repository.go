package repository

import (
	"github.com/Sirpyerre/fintech-backend/internal/dbconnection"
	"github.com/Sirpyerre/fintech-backend/internal/models"
)

type TransactionRepository struct {
	DBConnection *dbconnection.DBConnection
}

func NewTransactionRepository(dbConnection *dbconnection.DBConnection) *TransactionRepository {
	return &TransactionRepository{DBConnection: dbConnection}
}

func (r *TransactionRepository) StoreTransaction(records []models.Transaction) error {
	//TODO implement me
	panic("implement me")
}
