package http

import (
	"github.com/gorilla/mux"
)

func NewRouter(customerHandler *CustomerHandler, nationalityHandler *NationalityHandler) *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/customers", customerHandler.Create).Methods("POST")
	api.HandleFunc("/customers", customerHandler.GetAll).Methods("GET")
	api.HandleFunc("/customers/{id}", customerHandler.GetByID).Methods("GET")
	api.HandleFunc("/customers/{id}", customerHandler.Update).Methods("PUT")
	api.HandleFunc("/customers/{id}", customerHandler.Delete).Methods("DELETE")

	api.HandleFunc("/nationalities", nationalityHandler.GetAll).Methods("GET")

	return r
}
