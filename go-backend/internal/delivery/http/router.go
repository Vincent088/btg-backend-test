package http

import (
	"time"

	"github.com/gorilla/mux"
)

func NewRouter(customerHandler *CustomerHandler, nationalityHandler *NationalityHandler, apiKey string, allowedOrigin string) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.Use(CORSMiddleware(allowedOrigin))
	api.Use(RateLimitMiddleware(100, time.Minute))
	api.Use(APIKeyMiddleware(apiKey))

	api.HandleFunc("/customers", customerHandler.Create).Methods("POST", "OPTIONS")
	api.HandleFunc("/customers", customerHandler.GetAll).Methods("GET", "OPTIONS")
	api.HandleFunc("/customers/{id}", customerHandler.GetByID).Methods("GET", "OPTIONS")
	api.HandleFunc("/customers/{id}", customerHandler.Update).Methods("PUT", "OPTIONS")
	api.HandleFunc("/customers/{id}", customerHandler.Delete).Methods("DELETE", "OPTIONS")

	api.HandleFunc("/nationalities", nationalityHandler.GetAll).Methods("GET", "OPTIONS")

	return r
}
