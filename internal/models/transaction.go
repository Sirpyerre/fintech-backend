package models

import "time"

type Transaction struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Amount     float64   `json:"amount"`
	OccurredAt time.Time // datetime of the transaction
	BatchID    string    `db:"batch_id"`
}
