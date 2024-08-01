package service

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/naceto/tempstation/configs"
)

type Bootstrap struct {
	logger  *slog.Logger
	config  *configs.Config
	factory DependencyFactory
}

func NewBootstrap(logger *slog.Logger, config *configs.Config) *Bootstrap {
	return &Bootstrap{
		logger: logger,
		config: config,
	}
}

type Service struct {
	bs *Bootstrap
}

func NewService(bs *Bootstrap) *Service {
	return &Service{
		bs: bs,
	}
}

func (s *Service) Start(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(time.RFC1123)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("The time is: " + tm))
	}))

	wrapperMux := NewLogger(s.bs.logger, mux)
	return http.ListenAndServe(":8080", wrapperMux)
}

func (s *Service) Stop(ctx context.Context) error {
	return nil
}
