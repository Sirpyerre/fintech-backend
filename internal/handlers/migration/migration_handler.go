package migration

import (
	"errors"
	"mime/multipart"
	"net/http"
	"strconv"
	"strings"

	"github.com/Sirpyerre/fintech-backend/internal/services"
	"github.com/Sirpyerre/fintech-backend/pkg/common"
	"github.com/rs/zerolog"
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

// Migrate godoc
// @Summary Migrate transactions from a CSV file
// @Description Upload a CSV file to migrate transactions. Returns the number of skipped rows and a success message.
// @Tags Migration
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file to upload"
// @Success 200 {object} map[string]string "Migration successful"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /migrate [post]
func (h *MigrationHandler) Migrate(w http.ResponseWriter, r *http.Request) error {
	file, err := extractFile(r)
	if err != nil {
		return common.JSONError(w, http.StatusBadRequest, err.Error())
	}

	skipped, err := h.MigrationService.Migrate(r.Context(), file)
	if err != nil {
		return common.JSONError(w, http.StatusInternalServerError, err.Error())
	}

	return common.JSONSuccess(w, http.StatusOK, map[string]string{
		"skipped_rows": strconv.Itoa(skipped),
		"message":      "Migration successful",
	})
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
