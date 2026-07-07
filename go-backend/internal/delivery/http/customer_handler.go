package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Vincent088/btg-backend-test/go-backend/internal/entity"
	"github.com/Vincent088/btg-backend-test/go-backend/internal/usecase"
)

type CustomerHandler struct {
	usecase *usecase.CustomerUsecase
}

func NewCustomerHandler(u *usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{usecase: u}
}

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func respondError(w http.ResponseWriter, status int, message string) {
	respondJSON(w, status, map[string]string{"error": message})
}

func (h *CustomerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var c entity.Customer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if c.NationalityID == 0 || c.CstName == "" || c.CstEmail == "" {
		respondError(w, http.StatusBadRequest, "nationality_id, cst_name, and cst_email are required")
		return
	}

	id, err := h.usecase.CreateCustomerWithFamily(c)
	if err != nil {
		respondError(w, http.StatusBadRequest, "failed to create customer — check that nationality_id exists and all fields are valid")
		return
	}

	respondJSON(w, http.StatusCreated, map[string]int{"cst_id": id})
}

func (h *CustomerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	customers, err := h.usecase.GetAllCustomers()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, customers)
}

func (h *CustomerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	c, err := h.usecase.GetCustomerByID(id)
	if err != nil {
		respondError(w, http.StatusNotFound, "customer not found")
		return
	}
	respondJSON(w, http.StatusOK, c)
}

func (h *CustomerHandler) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	var c entity.Customer
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		respondError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	c.CstID = id

	if c.NationalityID == 0 || c.CstName == "" || c.CstEmail == "" {
		respondError(w, http.StatusBadRequest, "nationality_id, cst_name, and cst_email are required")
		return
	}

	if err := h.usecase.UpdateCustomerWithFamily(c); err != nil {
		respondError(w, http.StatusBadRequest, "failed to update customer — check that nationality_id exists and all fields are valid")
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "updated successfully"})
}

func (h *CustomerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		respondError(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.usecase.DeleteCustomer(id); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, map[string]string{"message": "deleted successfully"})
}
