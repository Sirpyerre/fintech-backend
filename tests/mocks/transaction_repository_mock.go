package mocks

import (
	"context"
	"github.com/Sirpyerre/fintech-backend/internal/models"
	"github.com/stretchr/testify/mock"
	"time"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) StoreTransaction(ctx context.Context, records []models.Transaction) error {
	args := m.Called(ctx, records)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) FindUserExists(ctx context.Context, userID int64) (bool, error) {
	args := m.Called(ctx, userID)
	return args.Bool(0), args.Error(1)
}

func (m *TransactionRepositoryMock) BalanceSummary(ctx context.Context, userID int64, from, to *time.Time) (float64, float64, float64, error) {
	args := m.Called(ctx, userID, from, to)
	return args.Get(0).(float64), args.Get(1).(float64), args.Get(2).(float64), args.Error(3)
}
