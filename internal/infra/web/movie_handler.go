package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebMovieHandler struct {
	MovieRepository entity.MovieRepositoryInterface
}

func NewMovieHandler(chooserRepository entity.MovieRepositoryInterface) *WebMovieHandler {
	return &WebMovieHandler{
		MovieRepository: chooserRepository,
	}
}

func (movieHandler *WebMovieHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecases.InputCreateMovieDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository)

	output, err := movieUseCase.Create(dto)
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

func (movieHandler *WebMovieHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository)

	movies, err := movieUseCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
