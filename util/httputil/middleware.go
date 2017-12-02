package httputil

import (
	"log"
	"net/http"
	"time"
)

var ApiKey string

type authorizationMiddleware struct {
	wrappedHandler http.Handler
}

func (a authorizationMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader != ApiKey {
		w.WriteHeader(http.StatusUnauthorized)
		return
	} else {
		a.wrappedHandler.ServeHTTP(w, r)
	}
}

func Authorize(h http.Handler) authorizationMiddleware {
	return authorizationMiddleware{h}
}

// Logs all requests
func Log(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		method := r.Method
		url := r.URL.String()
		log.Printf("%s: %s", method, url)
		h.ServeHTTP(w, r)
		log.Printf("Finished %s to %s: [%v]", method, url, time.Since(start))
	})
}
