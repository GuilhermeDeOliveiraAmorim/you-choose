package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebChooserHandler struct {
	ChooserRepository   entity.ChooserRepositoryInterface
	MovieListRepository entity.MovieListRepositoryInterface
}

func NewChooserHandler(chooserRepository entity.ChooserRepositoryInterface, movieListRepository entity.MovieListRepositoryInterface) *WebChooserHandler {
	return &WebChooserHandler{
		ChooserRepository:   chooserRepository,
		MovieListRepository: movieListRepository,
	}
}

func (chooserHandler *WebChooserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecases.InputCreateChooserDto
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository, chooserHandler.MovieListRepository)

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
	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository, chooserHandler.MovieListRepository)

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
	chooserId := r.URL.Query().Get("chooser_id")

	input := usecases.InputFindChooserDto{
		ID: chooserId,
	}

	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository, chooserHandler.MovieListRepository)

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
	chooserId := r.URL.Query().Get("chooser_id")

	input := usecases.InputDeleteChooserDto{
		ID: chooserId,
	}

	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository, chooserHandler.MovieListRepository)

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

func (chooserHandler *WebChooserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var input usecases.InputUpdateChooserDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository, chooserHandler.MovieListRepository)

	chooser, err := chooserUseCase.Update(input)
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

func (chooserHandler *WebChooserHandler) IsDeleted(w http.ResponseWriter, r *http.Request) {
	chooserId := r.URL.Query().Get("chooser_id")

	input := usecases.InputIsDeletedChooserDto{
		ID: chooserId,
	}

	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository, chooserHandler.MovieListRepository)

	chooser, err := chooserUseCase.IsDeleted(input)
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

func (chooserHandler *WebChooserHandler) CreateChooserAndMovieList(w http.ResponseWriter, r *http.Request) {
	var dto usecases.InputCreateChooserAndMovieListDto
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository, chooserHandler.MovieListRepository)

	output, err := chooserUseCase.CreateChooserAndMovieList(dto)
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

func (chooserHandler *WebChooserHandler) ChooserCreateMovieList(w http.ResponseWriter, r *http.Request) {
	var dto usecases.InputChooserCreateMovieListDto
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository, chooserHandler.MovieListRepository)

	output, err := chooserUseCase.ChooserCreateMovieList(dto)
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

func (chooserHandler *WebChooserHandler) FindAllChooserMovieLists(w http.ResponseWriter, r *http.Request) {
	chooserId := r.URL.Query().Get("chooser_id")

	input := usecases.InputFindAllChooserMovieListsDto{
		ChooserId: chooserId,
	}

	chooserUseCase := *usecases.NewChooserUseCase(chooserHandler.ChooserRepository, chooserHandler.MovieListRepository)

	allChooserMovieLists, err := chooserUseCase.FindAllChooserMovieLists(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(allChooserMovieLists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}