package services

import (
	"context"
	"encoding/csv"
	"github.com/Sirpyerre/fintech-backend/internal/models"
	"github.com/Sirpyerre/fintech-backend/internal/repository"
	"github.com/rs/zerolog"
	"io"
	"strconv"
	"time"
)

type MigrationService struct {
	transactionRepo repository.Transactioner
	logger          zerolog.Logger
}

func NewMigrationService(transactionRepo repository.Transactioner, logger zerolog.Logger) *MigrationService {
	return &MigrationService{
		transactionRepo: transactionRepo,
		logger:          logger,
	}
}

func (m MigrationService) Migrate(ctx context.Context, cvsFile io.Reader) error {
	transactions, err := parseCSV(cvsFile)
	if err != nil {
		m.logger.Printf("Error parsing CSV: %v", err)
		return err
	}

	if err := m.transactionRepo.StoreTransaction(ctx, transactions); err != nil {
		m.logger.Printf("Error storing transactions: %v", err)
		return err
	}

	return nil
}

func parseCSV(file io.Reader) ([]models.Transaction, error) {
	if file == nil {
		return nil, nil
	}

	reader := csv.NewReader(file)
	_, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		amount, err := parseAmount(record[2])
		if err != nil {
			return nil, err
		}

		date, err := parseDate(record[3])
		if err != nil {
			return nil, err
		}

		transaction := models.Transaction{
			ID:         record[0],
			UserID:     record[1],
			Amount:     amount,
			OccurredAt: date,
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func parseDate(dateStr string) (time.Time, error) {
	layout := time.RFC3339
	return time.Parse(layout, dateStr)
}

func parseAmount(amountStr string) (float64, error) {
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return 0, err
	}
	return amount, nil
}
