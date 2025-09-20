package services

import (
	"context"
	"github.com/Sirpyerre/fintech-backend/internal/repository"
	"io"
)

type MigrationService struct {
	transactionRepo repository.Transactioner
}

func (m MigrationService) Migrate(ctx context.Context, cvsFile io.Reader) error {
	//TODO implement me
	panic("implement me")
}

func NewMigrationService(transactionRepo repository.Transactioner) *MigrationService {
	return &MigrationService{transactionRepo: transactionRepo}
}
