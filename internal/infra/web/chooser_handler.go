package infra

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebChooserHandler struct {
	ChooserRepository entity.ChooserRepositoryInterface
}

func NewChooserHandler(chooserRepository entity.ChooserRepositoryInterface) *WebChooserHandler {
	return &WebChooserHandler{
		ChooserRepository: chooserRepository,
	}
}

func (h *WebChooserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecases.InputCreateChooserDto
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createChooser := *usecases.NewChooserUseCase(h.ChooserRepository)
	output, err := createChooser.Create(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
