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

	var from, to *time.Time
	if fromStr != "" {
		fromTime, err := time.Parse(time.RFC3339, fromStr)
		if err != nil {
			return common.JSONError(w, http.StatusBadRequest, "Invalid 'from' date format")
		}
		from = &fromTime
	}
	if toStr != "" {
		toTime, err := time.Parse(time.RFC3339, toStr)
		if err != nil {
			return common.JSONError(w, http.StatusBadRequest, "Invalid 'to' date format")
		}
		to = &toTime
	}

	balanceResponse, err := h.balanceService.Balance(r.Context(), int64(id), from, to)
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
