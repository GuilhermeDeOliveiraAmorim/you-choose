package web

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

	chooserUseCase := *usecases.NewChooserUseCase(h.ChooserRepository)
	// fmt.Println(dto)

	output, err := chooserUseCase.Create(dto)
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

func (h *WebChooserHandler) FindAll(w http.ResponseWriter, r *http.Request) {

	chooserUseCase := *usecases.NewChooserUseCase(h.ChooserRepository)

	choosers, err := chooserUseCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(choosers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebChooserHandler) Find(w http.ResponseWriter, r *http.Request) {
	chooserId := r.URL.Query().Get("id")

	// fmt.Println(chooserId)

	input := usecases.InputFindChooserDto{
		ID: chooserId,
	}

	chooserUseCase := *usecases.NewChooserUseCase(h.ChooserRepository)

	chooser, err := chooserUseCase.Find(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(chooser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
