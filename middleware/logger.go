package middleware

import (
	"log/slog"
	"net/http"

	"github.com/felixge/httpsnoop"
)

func Logger(next http.Handler) http.Handler {
	fn := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metrics := httpsnoop.CaptureMetrics(next, w, r)

		slog.Info("request",
			"host", r.RemoteAddr,
			"method", r.Method,
			"path", r.URL.Path,
			"protocol", r.Proto,
			"status", metrics.Code,
			"bytes", metrics.Written,
			"referer", r.Referer(),
			"user_agent", r.UserAgent(),
			"latency", metrics.Duration.Milliseconds(),
		)
	})

	return http.HandlerFunc(fn)
}
