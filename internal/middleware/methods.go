package middleware

import (
	"net/http"
	"slices"
)

// ValidateMethods is a middleware function that validates if a request
// is an allowed method for an url else return 400 Bad request
func ValidateMethods(methods ...string) Middleware {

	return func(f http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {

			if ok := slices.Contains(methods, r.Method); !ok {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			f(w, r)
		}
	}
}
