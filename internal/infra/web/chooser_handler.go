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

func (chooserHandler *WebChooserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecases.InputCreateChooserDto
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository)
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

func (chooserHandler *WebChooserHandler) FindAll(w http.ResponseWriter, r *http.Request) {

	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository)

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

func (chooserHandler *WebChooserHandler) Find(w http.ResponseWriter, r *http.Request) {
	chooserId := r.URL.Query().Get("id")

	input := usecases.InputFindChooserDto{
		ID: chooserId,
	}

	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository)

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

func (chooserHandler *WebChooserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	chooserId := r.URL.Query().Get("id")

	input := usecases.InputDeleteChooserDto{
		ID: chooserId,
	}

	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository)

	chooser, err := chooserUseCase.Delete(input)
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
