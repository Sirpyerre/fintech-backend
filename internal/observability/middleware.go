package observability

import (
	"github.com/rs/zerolog"
	"net/http"
	"time"
)

func LoggingMiddleware(logger zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Since(start)

			logger.Info().
				Str("method", r.Method).
				Str("url", r.URL.String()).
				Str("remote", r.RemoteAddr).
				Dur("duration", duration).
				Msg("request completed")
		})
	}
}
