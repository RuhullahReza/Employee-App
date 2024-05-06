package main

import (
	"os"
	"os/signal"
	"syscall"

	app "github.com/RuhullahReza/Employee-App/app"
	config "github.com/RuhullahReza/Employee-App/config"
	"github.com/RuhullahReza/Employee-App/pkg/logger"
)

func main() {
	logger.Init()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Log.Error(err, "failed to load config")
		os.Exit(1)
	}

	app := app.NewServer(cfg)

	q := make(chan os.Signal)
	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-q

		logger.Log.Info("Shutting down ....")
		if err := app.Shutdown(); err != nil {
			logger.Log.Error(err, "failed to shutdown gracefully")
		}
	}()

	logger.Log.Info("strating server")
	if err := app.Listen(cfg.AppHost); err != nil {
		logger.Log.Error(err, "failed to start server")
	}
}
