package featureflag

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"snapp-featureflag/internal/package/config"
	"syscall"
	"time"
)

type App struct {
	Config     *config.AppConfig
	httpServer *http.Server
}

func NewApp(config *config.AppConfig, httpServeMux *http.ServeMux) (*App, error) {
	if config == nil {
		return nil, errors.New("config can not be nil")
	}
	if httpServeMux == nil {
		return nil, errors.New("server can not be nil")
	}

	httpServer := &http.Server{
		Addr: fmt.Sprintf(
			"localhost:%d",
			config.Server.Port),
		Handler: httpServeMux,
	}

	app := &App{
		Config:     config,
		httpServer: httpServer,
	}

	return app, nil

}

func (a *App) createChannel() (chan os.Signal, func()) {
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt,
		syscall.SIGTERM, syscall.SIGINT)

	return stopCh, func() {
		close(stopCh)
	}
}

func (a *App) start(server *http.Server) {
	log.Println("application started")
	if err := server.ListenAndServe(); err != nil &&
		!errors.Is(err, http.ErrServerClosed) {
		panic(err)
	} else {
		log.Println("application stopped gracefully")
	}
}

func (a *App) shutdown(ctx context.Context, server *http.Server) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	} else {
		log.Println("application shutdown")
	}
}

func (a *App) Run() {
	go a.start(a.httpServer)

	stopCh, closeCh := a.createChannel()
	defer closeCh()
	log.Println("notified:", <-stopCh)

	a.shutdown(context.Background(), a.httpServer)
}
