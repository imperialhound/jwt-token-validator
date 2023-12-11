package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/iand/logfmtr"
	"github.com/imperialhound/friend-foe-api/internal/handlers"
	"github.com/imperialhound/friend-foe-api/internal/middleware"
	"github.com/imperialhound/friend-foe-api/internal/utils"
)

func main() {
	// Generate config file
	c := utils.NewConfig()

	// Initalize logger
	logfmtr.SetVerbosity(c.Verbosity)
	logger := utils.NewLogger()

	r := mux.NewRouter()
	r.HandleFunc("/", middleware.Chain(handlers.Sniff,
		middleware.LogRequestTime(logger),
		middleware.ValidateMethods("GET")))

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
	logger.Info("shutting down friend-foe server", "signal", sig)
}
