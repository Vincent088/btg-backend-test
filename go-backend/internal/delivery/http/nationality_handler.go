package http

import (
	"net/http"

	"github.com/Vincent088/btg-backend-test/go-backend/internal/usecase"
)

type NationalityHandler struct {
	usecase *usecase.NationalityUsecase
}

func NewNationalityHandler(u *usecase.NationalityUsecase) *NationalityHandler {
	return &NationalityHandler{usecase: u}
}

// GET /nationalities
func (h *NationalityHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.usecase.GetAllNationalities()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, list)
}
