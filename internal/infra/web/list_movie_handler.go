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
}

func NewMovieListHandler(movieListRepository entity.MovieListRepositoryInterface, chooserRepository entity.ChooserRepositoryInterface) *WebMovieListHandler {
	return &WebMovieListHandler{
		MovieListRepository: movieListRepository,
		ChooserRepository:   chooserRepository,
	}
}

func (h *WebMovieListHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecases.InputCreateMovieListDto
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieListUseCase := *usecases.NewMovieListUseCase(h.MovieListRepository, h.ChooserRepository)

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
	movieListUseCase := *usecases.NewMovieListUseCase(h.MovieListRepository, h.ChooserRepository)

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

	movieListUseCase := *usecases.NewMovieListUseCase(h.MovieListRepository, h.ChooserRepository)

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

func (h *WebMovieListHandler) AddChooserToMovieList(w http.ResponseWriter, r *http.Request) {
	chooserId := r.URL.Query().Get("chooser_id")
	movieListId := r.URL.Query().Get("movie_list_id")

	input := usecases.InputAddChooserToMovieListDto{
		ChooserId:   chooserId,
		MovieListId: movieListId,
	}

	movieListUseCase := *usecases.NewMovieListUseCase(h.MovieListRepository, h.ChooserRepository)

	chooserInMovieList, err := movieListUseCase.AddChooserToMovieList(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(chooserInMovieList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
