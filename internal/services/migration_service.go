package services

import (
	"context"
	"encoding/csv"
	"github.com/Sirpyerre/fintech-backend/internal/models"
	"github.com/Sirpyerre/fintech-backend/internal/repository"
	"github.com/rs/zerolog"
	"io"
	"strconv"
	"sync"
	"time"
)

type MigrationService struct {
	transactionRepo repository.Transactioner
	logger          zerolog.Logger
	workers         int
}

func NewMigrationService(transactionRepo repository.Transactioner, workers int, logger zerolog.Logger) *MigrationService {
	return &MigrationService{
		transactionRepo: transactionRepo,
		logger:          logger,
		workers:         workers,
	}
}

func (m *MigrationService) Migrate(ctx context.Context, cvsFile io.Reader) (int, error) {
	transactions, skipped, err := m.parseCSV(cvsFile)
	if err != nil {
		m.logger.Printf("Error parsing CSV: %v", err)
		return 0, err
	}

	if err := m.transactionRepo.StoreTransaction(ctx, transactions); err != nil {
		m.logger.Printf("Error storing transactions: %v", err)
		return 0, err
	}

	m.logger.Info().Int("skipped_rows", skipped).Msg("Migration completed")
	return skipped, nil
}

func (m *MigrationService) parseCSV(file io.Reader) ([]models.Transaction, int, error) {
	if file == nil {
		return nil, 0, nil
	}

	reader := csv.NewReader(file)
	_, err := reader.Read()
	if err != nil {
		return nil, 0, err
	}

	type parseResult struct {
		Tx  *models.Transaction
		Err error
	}

	jobs := make(chan []string)
	results := make(chan parseResult)
	var wg sync.WaitGroup

	skippedCh := make(chan int, m.workers)

	// workers
	for i := 0; i < m.workers; i++ {
		wg.Add(1)
		go func() {
			skipped := 0
			defer func() {
				skippedCh <- skipped
				wg.Done()
			}()
			for record := range jobs {
				tx, err := parseRecord(record)
				if err != nil {
					m.logger.Error().Err(err).Str("row", record[0]).Msg("Skipping row due to parse error")
					skipped++
					continue // skip this row
				}
				results <- parseResult{Tx: tx, Err: nil}
			}
		}()
	}

	go func() {
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				results <- parseResult{Tx: nil, Err: err}
				continue
			}
			jobs <- record
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
		close(skippedCh)
	}()

	var transactions []models.Transaction
	for res := range results {
		if res.Tx != nil {
			transactions = append(transactions, *res.Tx)
		}
	}

	skipped := 0
	for s := range skippedCh {
		skipped += s
	}

	return transactions, skipped, nil
}

func parseRecord(record []string) (*models.Transaction, error) {
	amount, err := parseAmount(record[2])
	if err != nil {
		return nil, err
	}

	date, err := parseDate(record[3])
	if err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:         record[0],
		UserID:     record[1],
		Amount:     amount,
		OccurredAt: date,
	}, nil
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
