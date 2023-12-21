package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain applies middleware functions to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
