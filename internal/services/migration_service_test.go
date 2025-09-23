package services

import (
	"context"
	"github.com/Sirpyerre/fintech-backend/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"strings"
	"testing"

	"github.com/rs/zerolog"
)

func TestMigrationService(t *testing.T) {
	logger := zerolog.Nop()
	ctx := context.Background()
	workerCount := 2

	t.Run("NewTransactionRepository initializes fields", func(t *testing.T) {
		repositoryMock := new(mocks.TransactionRepositoryMock)

		migrationService := NewMigrationService(repositoryMock, workerCount, logger)
		assert.IsType(t, &MigrationService{}, migrationService)
	})

	t.Run("Migrate fails with nil csv file", func(t *testing.T) {
		repositoryMock := new(mocks.TransactionRepositoryMock)
		repositoryMock.On("StoreTransaction", mock.Anything, mock.Anything).Return(assert.AnError)

		migrationService := NewMigrationService(repositoryMock, workerCount, logger)
		skipped, err := migrationService.Migrate(ctx, nil)
		assert.Error(t, err)
		assert.Equal(t, 0, skipped)
		assert.Equal(t, assert.AnError, err)
		repositoryMock.AssertExpectations(t)
	})

	// test fail with simulate a valid csv file input
	t.Run("Migrate fails with valid csv file", func(t *testing.T) {
		repositoryMock := new(mocks.TransactionRepositoryMock)
		repositoryMock.On("StoreTransaction", mock.Anything, mock.Anything).Return(assert.AnError)

		migrationService := NewMigrationService(repositoryMock, workerCount, logger)
		csvData := `id,user_id,amount,occurred_at
1,2,100.50,2023-10-01T10:00:00Z
2,1,200.75,2023-10-02T11:30:00Z
`
		skipped, err := migrationService.Migrate(ctx, io.NopCloser(strings.NewReader(csvData)))
		assert.Error(t, err)
		assert.Equal(t, 0, skipped)
		assert.Equal(t, assert.AnError, err)
		repositoryMock.AssertExpectations(t)
	})

	// test success with simulate a valid csv file input
	t.Run("Migrate succeeds with valid csv file", func(t *testing.T) {
		repositoryMock := new(mocks.TransactionRepositoryMock)
		repositoryMock.On("StoreTransaction", mock.Anything, mock.Anything).Return(nil)

		migrationService := NewMigrationService(repositoryMock, workerCount, logger)
		csvData := `id,user_id,amount,occurred_at
1,2,100.50,2023-10-01T10:00:00Z
2,1,200.75,2023-10-02T11:30:00Z
`
		skipped, err := migrationService.Migrate(ctx, io.NopCloser(strings.NewReader(csvData)))
		assert.NoError(t, err)
		assert.Equal(t, 0, skipped)
		repositoryMock.AssertExpectations(t)
	})
}
