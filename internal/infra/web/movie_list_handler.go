package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
)

type WebMovieListHandler struct {
	MovieListRepository entity.MovieListRepositoryInterface
}

func NewMovieListHandler(movieListRepository entity.MovieListRepositoryInterface) *WebMovieListHandler {
	return &WebMovieListHandler{
		MovieListRepository: movieListRepository,
	}
}

func (h *WebMovieListHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecases.InputCreateMovieListDto
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieListUseCase := *usecases.NewMovieListUseCase(h.MovieListRepository)
	// fmt.Println(dto)

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

func (h *WebMovieListHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	movieListUseCase := *usecases.NewMovieListUseCase(h.MovieListRepository)

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

func (h *WebMovieListHandler) Find(w http.ResponseWriter, r *http.Request) {
	movieListId := r.URL.Query().Get("id")

	input := usecases.InputFindMovieListDto{
		ID: movieListId,
	}

	movieListUseCase := *usecases.NewMovieListUseCase(h.MovieListRepository)

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
