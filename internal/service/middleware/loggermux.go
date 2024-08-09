package middleware

import (
	"log"
	"log/slog"
	"net/http"
	"time"
)

// Logger is a middleware handler that does request logging
type Logger struct {
	logger  *slog.Logger
	handler http.Handler
}

// ServeHTTP handles the request by passing it to the real
// handler and logging the request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(logger *slog.Logger, handlerToWrap http.Handler) *Logger {
	return &Logger{
		logger:  logger,
		handler: handlerToWrap,
	}
}
