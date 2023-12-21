package utils

import (
	"github.com/go-logr/logr"
	"github.com/iand/logfmtr"
)

func NewLogger() logr.Logger {
	opts := logfmtr.DefaultOptions()
	opts.Humanize = true
	opts.AddCaller = true
	return logfmtr.NewWithOptions(opts)
}
