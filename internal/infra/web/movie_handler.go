package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebMovieHandler struct {
	MovieRepository     entity.MovieRepositoryInterface
	ActorRepository     entity.ActorRepositoryInterface
	WriterRepository    entity.WriterRepositoryInterface
	DirectorRepository  entity.DirectorRepositoryInterface
	GenreRepository     entity.GenreRepositoryInterface
	MovieListRepository entity.MovieListRepositoryInterface
}

func NewMovieHandler(
	movieRepository entity.MovieRepositoryInterface,
	actorRepository entity.ActorRepositoryInterface,
	writerRepository entity.WriterRepositoryInterface,
	directorRepository entity.DirectorRepositoryInterface,
	genreRepository entity.GenreRepositoryInterface,
	movieListRepository entity.MovieListRepositoryInterface) *WebMovieHandler {
	return &WebMovieHandler{
		MovieRepository:     movieRepository,
		ActorRepository:     actorRepository,
		WriterRepository:    writerRepository,
		DirectorRepository:  directorRepository,
		GenreRepository:     genreRepository,
		MovieListRepository: movieListRepository,
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

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

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

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

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

func (movieHandler *WebMovieHandler) Update(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPut
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var input usecases.InputUpdateMovieDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

	movie, err := movieUseCase.Update(input)
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

func (movieHandler *WebMovieHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPatch
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieId := r.URL.Query().Get("movie_id")

	input := usecases.InputDeleteMovieDto{
		MovieId: movieId,
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

	movie, err := movieUseCase.Delete(input)
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

func (movieHandler *WebMovieHandler) IsDeleted(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieId := r.URL.Query().Get("movie_id")

	input := usecases.InputIsDeletedMovieDto{
		MovieId: movieId,
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

	movie, err := movieUseCase.IsDeleted(input)
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

func (movieHandler *WebMovieHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

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

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

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

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

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

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

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

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

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

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

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

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

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

func (movieHandler *WebMovieHandler) AddGenresToMovie(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputAddGenresToMovieDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

	output, err := movieUseCase.AddGenresToMovie(dto)
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

func (movieHandler *WebMovieHandler) FindMovieGenres(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieId := r.URL.Query().Get("movie_id")

	input := usecases.InputFindMovieGenresDto{
		MovieId: movieId,
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

	output, err := movieUseCase.FindMovieGenres(input)
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

func (movieHandler *WebMovieHandler) AddVoteToMovie(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPatch
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieId := r.URL.Query().Get("movie_id")

	input := usecases.InputAddVoteToMovieDto{
		MovieId: movieId,
	}

	movieUseCase := *usecases.NewMovieUseCase(movieHandler.MovieRepository, movieHandler.ActorRepository, movieHandler.WriterRepository, movieHandler.DirectorRepository, movieHandler.GenreRepository, movieHandler.MovieListRepository)

	output, err := movieUseCase.AddVoteToMovie(input)
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
