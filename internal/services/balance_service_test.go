package services

import (
	"context"
	"github.com/Sirpyerre/fintech-backend/tests/mocks"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestBalanceService(t *testing.T) {
	logger := zerolog.Nop()
	mockRepo := new(mocks.TransactionRepositoryMock)
	balanceService := NewBalanceService(mockRepo, logger)

	t.Run("User not found", func(t *testing.T) {
		userID := int64(1)
		mockRepo.On("FindUserExists", mock.Anything, userID).Return(false, nil).Once()

		_, err := balanceService.Balance(context.Background(), userID, nil, nil)
		assert.ErrorIs(t, err, ErrUserNotFound)
		mockRepo.AssertExpectations(t)
	})

	// invalid date range
	t.Run("Invalid date range", func(t *testing.T) {
		userID := int64(1)
		from := timePtr("2023-10-02T00:00:00Z")
		to := timePtr("2023-10-01T00:00:00Z")
		mockRepo.On("FindUserExists", mock.Anything, userID).Return(true, nil).Once()

		_, err := balanceService.Balance(context.Background(), userID, from, to)
		assert.ErrorIs(t, err, ErrInvalidDateRange)
		mockRepo.AssertExpectations(t)
	})

	// error getting BalanceSummary
	t.Run("Error getting BalanceSummary", func(t *testing.T) {
		userID := int64(1)
		mockRepo.On("FindUserExists", mock.Anything, userID).Return(true, nil).Once()
		mockRepo.On("BalanceSummary", mock.Anything, userID, (*time.Time)(nil), (*time.Time)(nil)).Return(0.0, 0.0, 0.0, assert.AnError).Once()

		_, err := balanceService.Balance(context.Background(), userID, nil, nil)
		assert.ErrorIs(t, err, assert.AnError)
		mockRepo.AssertExpectations(t)
	})

	// successful balance retrieval
	t.Run("Successful balance retrieval", func(t *testing.T) {
		userID := int64(1)
		mockRepo.On("FindUserExists", mock.Anything, userID).Return(true, nil).Once()
		mockRepo.On("BalanceSummary", mock.Anything, userID, (*time.Time)(nil), (*time.Time)(nil)).Return(150.567, 50.123, 100.444, nil).Once()

		balanceResp, err := balanceService.Balance(context.Background(), userID, nil, nil)
		assert.NoError(t, err)
		assert.Equal(t, 100.44, balanceResp.Balance)
		assert.Equal(t, 150.56, balanceResp.TotalCredits)
		assert.Equal(t, 50.12, balanceResp.TotalDebits)
		mockRepo.AssertExpectations(t)
	})
}

func timePtr(s string) *time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil
	}
	return &t
}
