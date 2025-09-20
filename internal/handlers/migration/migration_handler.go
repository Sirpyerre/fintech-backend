package migration

import (
	"github.com/Sirpyerre/fintech-backend/internal/services"
	"net/http"
)

type MigrationHandler struct {
	MigrationService services.Migrationer
}

func NewMigrationHandler(migrationService services.Migrationer) *MigrationHandler {
	return &MigrationHandler{MigrationService: migrationService}
}

func (h *MigrationHandler) Migrate(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseMultipartForm(10 << 20) // limit your max input length!
	if err != nil {
		http.Error(w, "Error parsing multipart form", http.StatusBadRequest)
		return err
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return err
	}
	defer file.Close()

	err = h.MigrationService.Migrate(r.Context(), file)
	if err != nil {
		http.Error(w, "Error migrating data", http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
