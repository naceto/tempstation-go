package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/naceto/tempstation/configs"
	"github.com/naceto/tempstation/internal/service"
	"github.com/naceto/tempstation/pkg/sdk"
)

func main() {
	logLevel := new(slog.LevelVar)
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel}))

	cfg, err := configs.New()
	if err != nil {
		logger.Error("error loading config", "error", err)
		panic(err)
	}

	level := sdk.IIf(cfg.Debug, slog.LevelDebug, slog.LevelInfo)
	logLevel.Set(level)

	bs := service.NewBootstrap(logger, cfg)
	service := service.NewService(bs)
	logger.Info("Starting service...")

	ctx := context.Background()
	if err := service.Start(ctx); err != nil {
		logger.Error("Service initialization failed", "error", err)
		panic(err)
	}

	logger.Info("Service started.")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	sig := <-stop

	logger.Info("Signal caught, stopping...", "signal", sig.String())
	service.Stop(ctx)
	logger.Info("Service stopped.")
}
