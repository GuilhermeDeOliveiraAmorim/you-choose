package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebMovieHandler struct {
	MovieRepository    entity.MovieRepositoryInterface
	ActorRepository    entity.ActorRepositoryInterface
	WriterRepository   entity.WriterRepositoryInterface
	DirectorRepository entity.DirectorRepositoryInterface
}

func NewMovieHandler(movieRepository entity.MovieRepositoryInterface, actorRepository entity.ActorRepositoryInterface, writerRepository entity.WriterRepositoryInterface, directorRepository entity.DirectorRepositoryInterface) *WebMovieHandler {
	return &WebMovieHandler{
		MovieRepository:    movieRepository,
		ActorRepository:    actorRepository,
		WriterRepository:   writerRepository,
		DirectorRepository: directorRepository,
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

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository)

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

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository)

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
		MovieId: movieId,
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository)

	movie, err := movieUseCase.Find(input)
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

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository)

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

func (movieHandler *WebMovieHandler) FindMovieActors(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieId := r.URL.Query().Get("movie_id")

	input := usecases.InputFindMovieActorsDto{
		MovieId: movieId,
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository)

	output, err := movieUseCase.FindMovieActors(input)
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

func (movieHandler *WebMovieHandler) AddWritersToMovie(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputAddWritersToMovieDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository)

	output, err := movieUseCase.AddWritersToMovie(dto)
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

func (movieHandler *WebMovieHandler) FindMovieWriters(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieId := r.URL.Query().Get("movie_id")

	input := usecases.InputFindMovieWritersDto{
		MovieId: movieId,
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository)

	output, err := movieUseCase.FindMovieWriters(input)
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

func (movieHandler *WebMovieHandler) AddDirectorsToMovie(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputAddDirectorsToMovieDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository)

	output, err := movieUseCase.AddDirectorsToMovie(dto)
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

func (movieHandler *WebMovieHandler) FindMovieDirectors(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieId := r.URL.Query().Get("movie_id")

	input := usecases.InputFindMovieDirectorsDto{
		MovieId: movieId,
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository)

	output, err := movieUseCase.FindMovieDirectors(input)
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
