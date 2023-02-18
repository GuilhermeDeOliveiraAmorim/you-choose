package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
)

type WebMovieListHandler struct {
	MovieListRepository entity.MovieListRepositoryInterface
	ChooserRepository   entity.ChooserRepositoryInterface
	MovieRepository     entity.MovieRepositoryInterface
	TagRepository       entity.TagRepositoryInterface
}

func NewMovieListHandler(
	movieListRepository entity.MovieListRepositoryInterface,
	chooserRepository entity.ChooserRepositoryInterface,
	movieRepository entity.MovieRepositoryInterface,
	tagRepository entity.TagRepositoryInterface) *WebMovieListHandler {
	return &WebMovieListHandler{
		MovieListRepository: movieListRepository,
		ChooserRepository:   chooserRepository,
		MovieRepository:     movieRepository,
	}
}

func (movieListHandler *WebMovieListHandler) Create(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputCreateMovieListDto
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieListUseCase := *usecases.NewMovieListUseCase(
		movieListHandler.MovieListRepository,
		movieListHandler.ChooserRepository,
		movieListHandler.MovieRepository,
		movieListHandler.TagRepository,
	)

	output, err := movieListUseCase.Create(dto)
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

func (movieListHandler *WebMovieListHandler) Find(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieListId := r.URL.Query().Get("movie_list_id")

	input := usecases.InputFindMovieListDto{
		MovieListId: movieListId,
	}

	movieListUseCase := *usecases.NewMovieListUseCase(
		movieListHandler.MovieListRepository,
		movieListHandler.ChooserRepository,
		movieListHandler.MovieRepository,
		movieListHandler.TagRepository,
	)

	movieList, err := movieListUseCase.Find(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(movieList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (movieListHandler *WebMovieListHandler) Update(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPut
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var input usecases.InputUpdateMovieListDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieListUseCase := *usecases.NewMovieListUseCase(
		movieListHandler.MovieListRepository,
		movieListHandler.ChooserRepository,
		movieListHandler.MovieRepository,
		movieListHandler.TagRepository,
	)

	movielist, err := movieListUseCase.Update(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(movielist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (movieListHandler *WebMovieListHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPatch
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieListId := r.URL.Query().Get("movieList_id")

	input := usecases.InputDeleteMovieListDto{
		MovieListId: movieListId,
	}

	movieListUseCase := *usecases.NewMovieListUseCase(
		movieListHandler.MovieListRepository,
		movieListHandler.ChooserRepository,
		movieListHandler.MovieRepository,
		movieListHandler.TagRepository,
	)

	movieList, err := movieListUseCase.Delete(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(movieList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (movieListHandler *WebMovieListHandler) IsDeleted(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieListId := r.URL.Query().Get("movieList_id")

	input := usecases.InputIsDeletedMovieListDto{
		MovieListId: movieListId,
	}

	movieListUseCase := *usecases.NewMovieListUseCase(
		movieListHandler.MovieListRepository,
		movieListHandler.ChooserRepository,
		movieListHandler.MovieRepository,
		movieListHandler.TagRepository,
	)

	movieList, err := movieListUseCase.IsDeleted(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(movieList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (movieListHandler *WebMovieListHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieListUseCase := *usecases.NewMovieListUseCase(
		movieListHandler.MovieListRepository,
		movieListHandler.ChooserRepository,
		movieListHandler.MovieRepository,
		movieListHandler.TagRepository,
	)

	movieLists, err := movieListUseCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(movieLists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (movieListHandler *WebMovieListHandler) FindMovieListMovies(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieListId := r.URL.Query().Get("movie_list_id")

	input := usecases.InputFindMovieListMoviesDto{
		MovieListId: movieListId,
	}

	movieListUseCase := *usecases.NewMovieListUseCase(
		movieListHandler.MovieListRepository,
		movieListHandler.ChooserRepository,
		movieListHandler.MovieRepository,
		movieListHandler.TagRepository,
	)

	output, err := movieListUseCase.FindMovieListMovies(input)
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

func (movieListHandler *WebMovieListHandler) FindMovieListChoosers(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	movieListId := r.URL.Query().Get("movie_list_id")

	input := usecases.InputFindMovieListChoosersDto{
		MovieListId: movieListId,
	}

	movieListUseCase := *usecases.NewMovieListUseCase(
		movieListHandler.MovieListRepository,
		movieListHandler.ChooserRepository,
		movieListHandler.MovieRepository,
		movieListHandler.TagRepository,
	)

	output, err := movieListUseCase.FindMovieListChoosers(input)
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
