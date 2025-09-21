package services

import (
	"context"
	"errors"
	"github.com/Sirpyerre/fintech-backend/internal/models"
	"github.com/Sirpyerre/fintech-backend/internal/repository"
	"github.com/rs/zerolog"
	"math"
	"time"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrInvalidDateRange = errors.New("invalid date range")
)

type BalanceService struct {
	transactionService repository.Transactioner
	logger             zerolog.Logger
}

func NewBalanceService(transactionService repository.Transactioner, logger zerolog.Logger) *BalanceService {
	return &BalanceService{
		transactionService: transactionService,
		logger:             logger,
	}
}

func (b BalanceService) Balance(ctx context.Context, userID int64, from, to *time.Time) (models.BalanceResponse, error) {
	exists, err := b.transactionService.FindUserExists(ctx, userID)
	if err != nil {
		return models.BalanceResponse{}, err
	}

	if !exists {
		return models.BalanceResponse{}, ErrUserNotFound
	}

	if from != nil && to != nil && from.After(*to) {
		return models.BalanceResponse{}, ErrInvalidDateRange
	}

	totalCredits, totalDebits, balance, err := b.transactionService.BalanceSummary(ctx, userID, from, to)
	if err != nil {
		return models.BalanceResponse{}, err
	}

	return models.BalanceResponse{
		Balance:      math.Floor(balance*100) / 100,
		TotalCredits: math.Floor(totalCredits*100) / 100,
		TotalDebits:  math.Floor(totalDebits*100) / 100,
	}, nil
}
