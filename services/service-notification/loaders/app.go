package loaders

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/kjurkovic/airtable/service/notification/config"
	"github.com/kjurkovic/airtable/service/notification/routes"
	"github.com/kjurkovic/airtable/service/notification/util"
	"github.com/rs/cors"
)

type App struct{}

func (app *App) Initialize() {
	logger := util.New()
	config, _ := config.Load()
	sm := mux.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   config.Server.AllowedOrigins,
		AllowCredentials: true,
		AllowedHeaders:   []string{"Authorization", "Content-Type", "Access-Control-Allow-Origin", "Origin"},
		AllowedMethods:   []string{"GET", "UPDATE", "PUT", "POST", "DELETE", "OPTIONS"},
		Debug:            config.Server.Debug,
	})

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.Port),
		Handler:      c.Handler(sm),
		IdleTimeout:  config.Server.Timeout.Idle * time.Second,
		ReadTimeout:  config.Server.Timeout.Read * time.Second,
		WriteTimeout: config.Server.Timeout.Write * time.Second,
	}

	routes := &routes.Routes{
		Router: sm,
		Logger: logger,
		Config: config,
	}

	routes.Initialize()

	go func() {
		logger.Info("Notification service listening on port %d", config.Server.Port)
		err := s.ListenAndServe()

		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	<-sigChannel

	logger.Info("Terminate received - shutting down gracefully")

	timeoutCtx, _ := context.WithTimeout(context.Background(), config.Server.Timeout.Shutdown*time.Second)
	s.Shutdown(timeoutCtx)
}
