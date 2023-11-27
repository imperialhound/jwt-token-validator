package middleware

import (
	"net/http"
	"time"

	"github.com/go-logr/logr"
)

// From https://gowebexamples.com/advanced-middleware/
func LogRequestTime(logger logr.Logger) Middleware {

	// Create new middlware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the Http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// log request time
			start := time.Now()
			logger.Info("request made", "request time", start)

			f(w, r)
		}
	}
}
