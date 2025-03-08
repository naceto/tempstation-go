package service

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/naceto/tempstation/configs"
	generic "github.com/naceto/tempstation/internal/generated/api/generic"
	sensors "github.com/naceto/tempstation/internal/generated/api/sensors"
	users "github.com/naceto/tempstation/internal/generated/api/users"
	"github.com/naceto/tempstation/internal/generated/db"
	"github.com/naceto/tempstation/internal/resources"
	"github.com/naceto/tempstation/internal/service/middleware"
	storage "github.com/naceto/tempstation/internal/storage/sqlc"
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

	queries := db.New(s.db)
	store := storage.NewStorage(queries)

	// Resources
	genericResource := resources.NewGeneric()
	sensorsResource := resources.NewSensors(s.bs.logger, store)
	usersResource := resources.NewUsers(s.bs.logger, store)

	root := http.NewServeMux()
	generic.HandlerFromMux(genericResource, root)

	api := http.NewServeMux()

	// Users
	us := users.NewStrictHandler(usersResource, nil)
	usersHandler := users.HandlerFromMux(us, api)
	uSwagger, err := users.GetSwagger()
	if err != nil {
		return err
	}

	// Sensors
	ss := sensors.NewStrictHandler(sensorsResource, nil)
	sensorsHandler := sensors.HandlerFromMux(ss, api)
	sSwagger, err := sensors.GetSwagger()
	if err != nil {
		return err
	}

	root.Handle("/api/swagger-ui/", http.StripPrefix("/api/swagger-ui", http.FileServerFS(web.Content)))
	root.Handle("/api/v1/users", http.StripPrefix("/api", strictMiddleware.OapiRequestValidator(uSwagger)(usersHandler)))
	root.Handle("/api/v1/sensors", http.StripPrefix("/api", strictMiddleware.OapiRequestValidator(sSwagger)(sensorsHandler)))

	logWrapper := middleware.NewLogger(s.bs.logger, root)
	return http.ListenAndServe(":8080", logWrapper)
}

func (s *Service) Stop(ctx context.Context) error {
	if s.db != nil {
		err := s.db.Close()
		if err != nil {
			s.bs.logger.Error("Closing DB connection failed", "error", err)
			return err
		}
	}

	return nil
}
