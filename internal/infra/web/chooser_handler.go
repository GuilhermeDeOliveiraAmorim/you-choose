package web

import (
	"encoding/json"
	"net/http"

	chooserCreateUseCase "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/application/usecases/chooser/create_chooser"

	chooserRepository "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/repository"
)

type WebChooserHandler struct {
	ChooserRepository chooserRepository.ChooserRepositoryInterface
}

func NewChooserHandler(chooserRepository chooserRepository.ChooserRepositoryInterface) *WebChooserHandler {
	return &WebChooserHandler{
		ChooserRepository: chooserRepository,
	}
}

func (h *WebChooserHandler) Create(w http.ResponseWriter, r *http.Request) {

	var dto chooserCreateUseCase.InputCreateChooserDto
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createChooser := *chooserCreateUseCase.NewCreateChooserUseCase(h.ChooserRepository)
	output, err := createChooser.Execute(dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// defer r.Body.Close()
}
