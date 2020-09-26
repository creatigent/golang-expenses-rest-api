package app

import (
	"context"
	"fmt"
	"github.com/steevehook/expenses-rest-api/config"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/steevehook/expenses-rest-api/controllers"
	"github.com/steevehook/expenses-rest-api/logging"
)

type App struct {
	stopOnce sync.Once
	Server   *http.Server
}

// Init initializes the application
func Init(configPath string) (*App, error) {
	configManager, err := config.Init(configPath)
	if err != nil {
		return nil, fmt.Errorf("could not initialize app config: %v", err)
	}

	if err := logging.Init(configManager); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("could not initialize logger: %v", err)
	}

	app := &App{
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      controllers.NewRouter(),
			ReadTimeout:  200 * time.Microsecond,
			WriteTimeout: 200 * time.Microsecond,
			ErrorLog:     logging.HTTPServerLogger(),
		},
	}
	return app, nil
}

// Start starts the application
func (a *App) Start() error {
	logging.Logger.Info(
		"http server is ready to handle requests",
		zap.String("listen", ":8080"),
	)

	err := a.Server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

// Stop shuts down the http server
func (a *App) Stop() {
	a.stopOnce.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()

		logging.Logger.Info("shutting down the http server")
		if err := a.Server.Shutdown(ctx); err != nil {
			logging.Logger.Error("error on server shutdown", zap.Error(err))
		}

		logging.Logger.Info("http server was shut down")
	})
}

// Stopper represents app stop feature
type Stopper interface {
	Stop()
}

// ListenToSignals listens for any incoming termination signals and shuts down the application
func ListenToSignals(signals []os.Signal, apps ...Stopper) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)

	<-c
	for _, a := range apps {
		a.Stop()
	}

	os.Exit(0)
}
