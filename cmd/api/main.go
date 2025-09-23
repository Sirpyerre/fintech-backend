package main

import (
	"context"
	"log"
	"net/http"

	"github.com/Sirpyerre/fintech-backend/internal/api"
	"github.com/Sirpyerre/fintech-backend/internal/config"
	"github.com/Sirpyerre/fintech-backend/internal/dbconnection"
	"github.com/Sirpyerre/fintech-backend/internal/handlers/balance"
	"github.com/Sirpyerre/fintech-backend/internal/handlers/migration"
	"github.com/Sirpyerre/fintech-backend/internal/observability"
	"github.com/Sirpyerre/fintech-backend/internal/repository"
	"github.com/Sirpyerre/fintech-backend/internal/services"

	_ "github.com/Sirpyerre/fintech-backend/docs"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfiguration(ctx)

	// logger setup
	logger := observability.InitLogger(cfg.LogLevel)
	logger.Info().Msg("Logger initialized")

	dbConn, err := dbconnection.NewDBConnection(ctx, cfg.DBConfig.DatabaseURL)
	if err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to the database")
	}

	// Initialize repositories, services, and handlers
	transactionRepo := repository.NewTransactionRepository(dbConn, logger)
	transactionService := services.NewMigrationService(transactionRepo, cfg.WorkerCount, logger)
	migrationHandler := migration.NewMigrationHandler(transactionService, logger)

	balanceService := services.NewBalanceService(transactionRepo, logger)
	balanceHandler := balance.NewBalanceHandler(balanceService, logger)

	r := api.NewRouter(migrationHandler, balanceHandler, logger)

	// Start the server
	logger.Info().Msgf("Starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
