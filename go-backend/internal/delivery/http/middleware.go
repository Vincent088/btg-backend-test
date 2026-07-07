package http

import (
	"net/http"
	"sync"
	"time"
)

func APIKeyMiddleware(apiKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get("X-API-Key")
			if key == "" || key != apiKey {
				respondError(w, http.StatusUnauthorized, "missing or invalid API key")
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func CORSMiddleware(allowedOrigin string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func RateLimitMiddleware(maxRequests int, window time.Duration) func(http.Handler) http.Handler {
	type client struct {
		count   int
		resetAt time.Time
	}

	var mu sync.Mutex
	clients := make(map[string]*client)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr

			mu.Lock()
			c, exists := clients[ip]
			if !exists || time.Now().After(c.resetAt) {
				c = &client{count: 0, resetAt: time.Now().Add(window)}
				clients[ip] = c
			}
			c.count++
			blocked := c.count > maxRequests
			mu.Unlock()

			if blocked {
				respondError(w, http.StatusTooManyRequests, "rate limit exceeded, try again later")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
