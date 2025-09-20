package main

import (
	"context"
	"github.com/Sirpyerre/fintech-backend/internal/config"
	"github.com/Sirpyerre/fintech-backend/internal/dbconnection"
	"github.com/Sirpyerre/fintech-backend/internal/handlers/health"
	"github.com/Sirpyerre/fintech-backend/internal/handlers/migration"
	"github.com/Sirpyerre/fintech-backend/internal/repository"
	"github.com/Sirpyerre/fintech-backend/internal/services"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	ctx := context.Background()
	cfg := config.NewConfiguration(ctx)
	dbConn, err := dbconnection.NewDBConnection(ctx, cfg.DBConfig.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	transactionRepo := repository.NewTransactionRepository(dbConn)
	transactionService := services.NewMigrationService(transactionRepo)
	migrationHandler := migration.NewMigrationHandler(transactionService)

	r := chi.NewRouter()
	r.Post("/migrate", func(w http.ResponseWriter, r *http.Request) {
		if err := migrationHandler.Migrate(w, r); err != nil {
			http.Error(w, "Error migrating data", http.StatusInternalServerError)
		}
	})
	r.Get("/health", health.HealthHandler)

	log.Printf("Starting server on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
