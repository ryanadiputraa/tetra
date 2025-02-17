package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Limit(100), 100)

func ThrottleMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("too many request"))
			return
		}
		h.ServeHTTP(w, r)
	})
}
