package balance

import (
	"errors"
	"github.com/Sirpyerre/fintech-backend/internal/services"
	"github.com/Sirpyerre/fintech-backend/pkg/common"
	"github.com/rs/zerolog"
	"net/http"
	"strconv"
	"time"
)

type BalanceHandler struct {
	balanceService services.Balancer
	logger         zerolog.Logger
}

func NewBalanceHandler(balanceService services.Balancer, logger zerolog.Logger) *BalanceHandler {
	return &BalanceHandler{
		balanceService: balanceService,
		logger:         logger,
	}
}

// GetBalance godoc
// @Summary      Get user balance
// @Description  Retrieves the balance for a specific user, optionally filtered by date range.
// @Tags balance
// @Produce json
// @Param user_id path int true "User ID"
// @Param from query string false "Start date in RFC3339 format"
// @Param to query string false "End date in RFC3339 format"
// @Success 200 {object} models.BalanceResponse
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "User not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /users/{user_id}/balance [get]
func (h *BalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) error {
	userID := r.PathValue("user_id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return common.JSONError(w, http.StatusBadRequest, err.Error())
	}

	// get from and to from query parameters
	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")
	dateRange, err := h.datetimeRange(fromStr, toStr)
	if err != nil {
		return common.JSONError(w, http.StatusBadRequest, "Invalid date format. Use RFC3339 format.")
	}

	balanceResponse, err := h.balanceService.Balance(r.Context(), int64(id), dateRange["from"], dateRange["to"])
	if err != nil {
		if errors.Is(err, services.ErrUserNotFound) {
			return common.JSONError(w, http.StatusNotFound, err.Error())
		}
		if errors.Is(err, services.ErrInvalidDateRange) {
			return common.JSONError(w, http.StatusBadRequest, err.Error())
		}
		h.logger.Error().Err(err).Msg("Failed to get balance")
		return common.JSONError(w, http.StatusInternalServerError, "Internal server error")
	}

	return common.JSONSuccess(w, http.StatusOK, balanceResponse)
}

func (h *BalanceHandler) datetimeRange(fromStr, toStr string) (map[string]*time.Time, error) {
	from, err := h.parseDate(fromStr)
	if err != nil {
		return nil, err
	}

	to, err := h.parseDate(toStr)
	if err != nil {
		return nil, err
	}

	return map[string]*time.Time{
		"from": from,
		"to":   to,
	}, nil
}

func (h *BalanceHandler) parseDate(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}
	// Try RFC3339
	t, err := time.Parse(time.RFC3339, dateStr)
	if err == nil {
		return &t, nil
	}
	// Try YYYY-MM-DD
	t, err = time.Parse("2006-01-02", dateStr)
	if err == nil {
		return &t, nil
	}
	h.logger.Error().Err(err).Msg("Failed to parse date")
	return nil, err
}
