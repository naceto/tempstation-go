package service

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/naceto/tempstation/configs"
	generic "github.com/naceto/tempstation/internal/generated/api/generic"
	sensors "github.com/naceto/tempstation/internal/generated/api/sensors"
	"github.com/naceto/tempstation/internal/handlers"
	"github.com/naceto/tempstation/internal/service/middleware"
	"github.com/naceto/tempstation/web"
	strictMiddleware "github.com/oapi-codegen/nethttp-middleware"
)

type Bootstrap struct {
	logger  *slog.Logger
	cfg     *configs.Config
	factory DependencyFactory
}

func NewBootstrap(logger *slog.Logger, config *configs.Config, factory DependencyFactory) *Bootstrap {
	return &Bootstrap{
		logger:  logger,
		cfg:     config,
		factory: factory,
	}
}

type Service struct {
	bs *Bootstrap
	db DB
}

func NewService(bs *Bootstrap) *Service {
	return &Service{
		bs: bs,
	}
}

func (s *Service) Start(ctx context.Context) error {
	var err error
	s.db, err = s.bs.factory.GetDB(ctx, s.bs.cfg)
	if err != nil {
		return err
	}

	root := http.NewServeMux()
	generic.HandlerFromMux(handlers.NewGeneric(), root)

	api := http.NewServeMux()
	ss := sensors.NewStrictHandler(handlers.NewSensors(), nil)
	sensorsHandler := sensors.HandlerFromMux(ss, api)
	sSwagger, err := sensors.GetSwagger()
	if err != nil {
		return err
	}

	root.Handle("/api/v1/sensors", http.StripPrefix("/api", strictMiddleware.OapiRequestValidator(sSwagger)(sensorsHandler)))
	root.Handle("/api/swagger-ui/", http.StripPrefix("/api/swagger-ui", http.FileServerFS(web.Content)))

	logWrapper := middleware.NewLogger(s.bs.logger, root)
	return http.ListenAndServe(":8080", logWrapper)
}

func (s *Service) Stop(ctx context.Context) error {
	if s.db != nil {
		err := s.db.Close()
		if err != nil {
			s.bs.logger.Error("Closing DB connection failed", "error", err)
		}
	}

	return nil
}
