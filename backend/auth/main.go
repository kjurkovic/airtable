package main

import (
	"auth/config"
	"auth/database"
	"auth/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	muxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	_, b, _, _  = runtime.Caller(0)
	projectRoot = filepath.Dir(b)
)

func main() {
	logger := log.New(os.Stdout, "[auth-api] ", log.LstdFlags)
	config, err := config.Load(projectRoot + "/config.yml")

	if err != nil {
		logger.Fatal(err)
	}

	// initialize database
	db := database.Get(&config.Database, logger)
	gormDb, _ := db.Db.DB()
	defer gormDb.Close()

	sm := mux.NewRouter()

	router := &router.Router{}
	router.Setup(logger, config, sm)

	// CORS
	corsHandler := muxHandlers.CORS(
		muxHandlers.AllowedOrigins(
			config.Server.AllowedOrigins,
		),
	)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Server.Port),
		Handler:      corsHandler(sm),
		IdleTimeout:  config.Server.Timeout.Idle * time.Second,
		ReadTimeout:  config.Server.Timeout.Read * time.Second,
		WriteTimeout: config.Server.Timeout.Write * time.Second,
	}

	go func() {
		err := s.ListenAndServe()

		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, syscall.SIGTERM)

	sig := <-sigChannel
	logger.Print("Terminate received - shutting down gracefully: ", sig)

	timeoutCtx, _ := context.WithTimeout(context.Background(), config.Server.Timeout.Shutdown*time.Second)
	s.Shutdown(timeoutCtx)
}
