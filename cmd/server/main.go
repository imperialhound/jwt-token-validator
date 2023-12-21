package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/iand/logfmtr"
	"github.com/imperialhound/friend-foe-api/internal/handlers"
	"github.com/imperialhound/friend-foe-api/internal/utils"
)

func main() {
	// Generate config file
	c := utils.NewConfig()

	// Initalize context with cancel
	ctx, cancel := context.WithCancel(context.Background())

	// Initalize logger
	logfmtr.SetVerbosity(c.Verbosity)
	logger := utils.NewLogger()

	// Add router and validator handler
	r := mux.NewRouter()

	validatorHandler := handlers.NewValidatorHandler(ctx, logger, c.AuthServer)
	r.HandleFunc("/", validatorHandler.ValidateToken)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", c.Port),
		Handler: r,
	}

	go func() {
		logger.Info("starting friend-foe server", "port", c.Port)

		err := server.ListenAndServe()
		if err != nil {
			logger.Error(err, "failed to start friend-foe server")
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	cancel()
	logger.Info("shutting down friend-foe server", "signal", sig)
}
