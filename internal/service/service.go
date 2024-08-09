package service

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/naceto/tempstation/configs"
	generic "github.com/naceto/tempstation/internal/generated/api/generic"
	sensors "github.com/naceto/tempstation/internal/generated/api/sensors"
	"github.com/naceto/tempstation/internal/resources"
	"github.com/naceto/tempstation/internal/service/middleware"
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
	r := resources.NewGenericResource()
	_ = generic.HandlerFromMux(r, mux)

	api := http.NewServeMux()
	sen := resources.NewSensorsResource()
	_ = sensors.HandlerFromMux(sen, api)
	mux.Handle("/api/", http.StripPrefix("/api", api))

	wrapperMux := middleware.NewLogger(s.bs.logger, mux)
	return http.ListenAndServe(":8080", wrapperMux)
}

func (s *Service) Stop(ctx context.Context) error {
	return nil
}
