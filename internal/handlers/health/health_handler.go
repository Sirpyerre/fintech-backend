package health

import (
	"net/http"
)

// HealthHandler responds with a simple "OK" message to indicate the service is running.
// @Summary Health Check
// @Tags Health
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /health [get]
func HealthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("OK"))
}
