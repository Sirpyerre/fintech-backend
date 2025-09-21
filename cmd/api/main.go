package main

import (
	"context"
	"github.com/Sirpyerre/fintech-backend/internal/handlers/balance"
	"log"
	"net/http"

	"github.com/Sirpyerre/fintech-backend/internal/config"
	"github.com/Sirpyerre/fintech-backend/internal/dbconnection"
	"github.com/Sirpyerre/fintech-backend/internal/handlers/health"
	"github.com/Sirpyerre/fintech-backend/internal/handlers/migration"
	"github.com/Sirpyerre/fintech-backend/internal/observability"
	"github.com/Sirpyerre/fintech-backend/internal/repository"
	"github.com/Sirpyerre/fintech-backend/internal/services"

	_ "github.com/Sirpyerre/fintech-backend/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfiguration(ctx)
	dbConn, err := dbconnection.NewDBConnection(ctx, cfg.DBConfig.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// logger setup
	logger := observability.InitLogger(cfg.LogLevel)
	logger.Info().Msg("Logger initialized")

	// Initialize repositories, services, and handlers
	transactionRepo := repository.NewTransactionRepository(dbConn, logger)
	transactionService := services.NewMigrationService(transactionRepo, logger)
	migrationHandler := migration.NewMigrationHandler(transactionService, logger)

	balanceService := services.NewBalanceService(transactionRepo, logger)
	balanceHandler := balance.NewBalanceHandler(balanceService, logger)

	r := chi.NewRouter()
	r.Use(observability.LoggingMiddleware(logger))
	//routes
	r.Post("/migrate", func(w http.ResponseWriter, r *http.Request) {
		if err := migrationHandler.Migrate(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Get("/users/{user_id}/balance", func(w http.ResponseWriter, r *http.Request) {
		if err := balanceHandler.GetBalance(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	r.Get("/health", health.HealthHandler)
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	// Start the server
	logger.Info().Msgf("Starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
