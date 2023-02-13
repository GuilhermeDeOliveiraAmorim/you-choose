package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebMovieHandler struct {
	MovieRepository entity.MovieRepositoryInterface
	ActorRepository entity.ActorRepositoryInterface
}

func NewMovieHandler(movieRepository entity.MovieRepositoryInterface, actorRepository entity.ActorRepositoryInterface) *WebMovieHandler {
	return &WebMovieHandler{
		MovieRepository: movieRepository,
		ActorRepository: actorRepository,
	}
}

func (movieHandler *WebMovieHandler) Create(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputCreateMovieDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository)

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
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository)

	movies, err := movieUseCase.MovieRepository.FindAll()
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

func (movieHandler *WebMovieHandler) Find(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieId := r.URL.Query().Get("movie_id")

	input := usecases.InputFindMovieDto{
		ID: movieId,
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository)

	movie, err := movieUseCase.MovieRepository.Find(input.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (movieHandler *WebMovieHandler) AddActorsToMovie(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputAddActorsToMovieDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository)

	output, err := movieUseCase.AddActorsToMovie(dto)
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
