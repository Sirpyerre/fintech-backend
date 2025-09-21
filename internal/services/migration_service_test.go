package services

import (
	"context"
	"github.com/Sirpyerre/fintech-backend/internal/models"
	"github.com/Sirpyerre/fintech-backend/internal/repository"
	"github.com/rs/zerolog"
	"io"
	"reflect"
	"testing"
	"time"
)

func TestMigrationService_Migrate(t *testing.T) {
	type fields struct {
		transactionRepo repository.Transactioner
		logger          zerolog.Logger
	}
	type args struct {
		ctx     context.Context
		cvsFile io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MigrationService{
				transactionRepo: tt.fields.transactionRepo,
				logger:          tt.fields.logger,
			}
			if err := m.Migrate(tt.args.ctx, tt.args.cvsFile); (err != nil) != tt.wantErr {
				t.Errorf("Migrate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewMigrationService(t *testing.T) {
	type args struct {
		transactionRepo repository.Transactioner
		logger          zerolog.Logger
	}
	tests := []struct {
		name string
		args args
		want *MigrationService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMigrationService(tt.args.transactionRepo, tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMigrationService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseAmount(t *testing.T) {
	type args struct {
		amountStr string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseAmount(tt.args.amountStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseAmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseAmount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseCSV(t *testing.T) {
	type args struct {
		file io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Transaction
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCSV(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCSV() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseDate(t *testing.T) {
	type args struct {
		dateStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDate(tt.args.dateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
