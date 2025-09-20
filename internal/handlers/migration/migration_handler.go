package migration

import (
	"errors"
	"github.com/Sirpyerre/fintech-backend/internal/services"
	"github.com/Sirpyerre/fintech-backend/pkg/common"
	"github.com/rs/zerolog"
	"mime/multipart"
	"net/http"
	"strings"
)

type MigrationHandler struct {
	MigrationService services.Migrationer
	logger           zerolog.Logger
}

func NewMigrationHandler(migrationService services.Migrationer, logger zerolog.Logger) *MigrationHandler {
	return &MigrationHandler{
		MigrationService: migrationService,
		logger:           logger,
	}
}

func (h *MigrationHandler) Migrate(w http.ResponseWriter, r *http.Request) error {
	file, err := extractFile(r)
	if err != nil {
		return common.JSONError(w, http.StatusBadRequest, err.Error())
	}

	err = h.MigrationService.Migrate(r.Context(), file)
	if err != nil {
		return common.JSONError(w, http.StatusInternalServerError, err.Error())
	}

	return common.JSONSuccess(w, http.StatusOK, map[string]string{"message": "Migration successful"})
}

func extractFile(r *http.Request) (file multipart.File, err error) {
	if r.Header.Get("Content-Type") == "" || !strings.HasPrefix(r.Header.Get("Content-Type"), "multipart/form-data") {
		return nil, errors.New("Content-Type must be multipart/form-data")
	}

	err = r.ParseMultipartForm(10 << 20) // limit your max input length!
	if err != nil {
		return nil, err
	}

	file, _, err = r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return file, nil
}
