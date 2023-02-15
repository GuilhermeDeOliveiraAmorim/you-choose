package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebGenreHandler struct {
	GenreRepository entity.GenreRepositoryInterface
	MovieRepository entity.MovieRepositoryInterface
}

func NewGenreHandler(genreRepository entity.GenreRepositoryInterface, movieRepository entity.MovieRepositoryInterface) *WebGenreHandler {
	return &WebGenreHandler{
		GenreRepository: genreRepository,
		MovieRepository: movieRepository,
	}
}

func (genreHandler *WebGenreHandler) Create(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputCreateGenreDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	genreUseCase := *usecases.NewGenreUseCase(genreHandler.GenreRepository, genreHandler.MovieRepository)

	output, err := genreUseCase.Create(dto)
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

func (genreHandler *WebGenreHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	genreUseCase := *usecases.NewGenreUseCase(genreHandler.GenreRepository, genreHandler.MovieRepository)

	genres, err := genreUseCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(genres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (genreHandler *WebGenreHandler) Find(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	genreId := r.URL.Query().Get("genre_id")

	input := usecases.InputFindGenreDto{
		ID: genreId,
	}

	genreUseCase := *usecases.NewGenreUseCase(genreHandler.GenreRepository, genreHandler.MovieRepository)

	genre, err := genreUseCase.Find(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
