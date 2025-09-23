package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/Sirpyerre/fintech-backend/internal/handlers/balance"
	"github.com/Sirpyerre/fintech-backend/internal/handlers/health"
	"github.com/Sirpyerre/fintech-backend/internal/handlers/migration"
	"github.com/Sirpyerre/fintech-backend/internal/observability"
)

func NewRouter(migrationHandler *migration.MigrationHandler, balanceHandler *balance.BalanceHandler, logger zerolog.Logger) *chi.Mux {
	r := chi.NewRouter()
	r.Use(observability.LoggingMiddleware(logger))

	// Migration route
	r.Post("/migrate", func(w http.ResponseWriter, r *http.Request) {
		if err := migrationHandler.Migrate(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Balance route
	r.Get("/users/{user_id}/balance", func(w http.ResponseWriter, r *http.Request) {
		if err := balanceHandler.GetBalance(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Health route
	r.Get("/health", health.HealthHandler)

	// Swagger docs
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	return r
}
