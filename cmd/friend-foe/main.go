package main

import (
	"fmt"
	"net/http"

	"github.com/alecthomas/kong"
	"github.com/go-logr/logr"
	"github.com/gorilla/mux"
	"github.com/iand/logfmtr"
	"github.com/imperialhound/friend-foe-api/internal/handlers"
	"github.com/imperialhound/friend-foe-api/internal/middleware"
)

type CLI struct {
	AuthServer string `help:"URL for auth server to validate JWTs." default:"localhost:9000"`
	Verbosity  int    `help:"Logging verbosity" default:"1"`
	Port       string `help:"Server listening port" default:"8080"`
}

func (c *CLI) Run() error {
	logfmtr.SetVerbosity(c.Verbosity)
	logger := newLogger()

	r := mux.NewRouter()
	r.HandleFunc("/", middleware.Chain(handlers.Sniff,
		middleware.LogRequestTime(logger),
		middleware.ValidateMethods("GET")))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", c.Port),
		Handler: r,
	}

	logger.Info("starting friend-foe server", "port", c.Port)
	return server.ListenAndServe()
}

func newLogger() logr.Logger {
	opts := logfmtr.DefaultOptions()
	opts.Humanize = true
	opts.AddCaller = true
	return logfmtr.NewWithOptions(opts)
}

func main() {
	ff := kong.Parse(&CLI{}, kong.Description(`An API to determine if my dog can sense 
if you are a friend or foe based on your smell (JWT)`))

	ff.FatalIfErrorf(ff.Run())
}
