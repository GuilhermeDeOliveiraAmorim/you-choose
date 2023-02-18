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
	MovieRepository     entity.MovieRepositoryInterface
	TagRepository       entity.TagRepositoryInterface
}

func NewChooserHandler(
	chooserRepository entity.ChooserRepositoryInterface,
	movieListRepository entity.MovieListRepositoryInterface,
	movieRepository entity.MovieRepositoryInterface,
	tagRepository entity.TagRepositoryInterface) *WebChooserHandler {
	return &WebChooserHandler{
		ChooserRepository:   chooserRepository,
		MovieListRepository: movieListRepository,
		MovieRepository:     movieRepository,
		TagRepository:       tagRepository,
	}
}

func (chooserHandler *WebChooserHandler) Create(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputCreateChooserDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chooserUseCase := *usecases.NewChooserUseCase(
		chooserHandler.ChooserRepository,
		chooserHandler.MovieListRepository,
		chooserHandler.MovieRepository,
		chooserHandler.TagRepository)

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

func (chooserHandler *WebChooserHandler) Find(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	chooserId := r.URL.Query().Get("chooser_id")

	input := usecases.InputFindChooserDto{
		ChooserId: chooserId,
	}

	chooserUseCase := *usecases.NewChooserUseCase(
		chooserHandler.ChooserRepository,
		chooserHandler.MovieListRepository,
		chooserHandler.MovieRepository,
		chooserHandler.TagRepository)

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

func (chooserHandler *WebChooserHandler) Update(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPut
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var input usecases.InputUpdateChooserDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chooserUseCase := *usecases.NewChooserUseCase(
		chooserHandler.ChooserRepository,
		chooserHandler.MovieListRepository,
		chooserHandler.MovieRepository,
		chooserHandler.TagRepository)

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

func (chooserHandler *WebChooserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodDelete
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	chooserId := r.URL.Query().Get("chooser_id")

	input := usecases.InputDeleteChooserDto{
		ChooserId: chooserId,
	}

	chooserUseCase := *usecases.NewChooserUseCase(
		chooserHandler.ChooserRepository,
		chooserHandler.MovieListRepository,
		chooserHandler.MovieRepository,
		chooserHandler.TagRepository)

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

func (chooserHandler *WebChooserHandler) IsDeleted(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	chooserId := r.URL.Query().Get("chooser_id")

	input := usecases.InputIsDeletedChooserDto{
		ChooserId: chooserId,
	}

	chooserUseCase := *usecases.NewChooserUseCase(
		chooserHandler.ChooserRepository,
		chooserHandler.MovieListRepository,
		chooserHandler.MovieRepository,
		chooserHandler.TagRepository)

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

func (chooserHandler *WebChooserHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	chooserUseCase := *usecases.NewChooserUseCase(
		chooserHandler.ChooserRepository,
		chooserHandler.MovieListRepository,
		chooserHandler.MovieRepository,
		chooserHandler.TagRepository)

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

func (chooserHandler *WebChooserHandler) AddMoviesToMovieList(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputAddMoviesToMovieListDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chooserUseCase := *usecases.NewChooserUseCase(
		chooserHandler.ChooserRepository,
		chooserHandler.MovieListRepository,
		chooserHandler.MovieRepository,
		chooserHandler.TagRepository)

	output, err := chooserUseCase.AddMoviesToMovieList(dto)
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

func (chooserHandler *WebChooserHandler) AddChoosersToMovieList(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputAddChoosersToMovieListDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chooserUseCase := *usecases.NewChooserUseCase(
		chooserHandler.ChooserRepository,
		chooserHandler.MovieListRepository,
		chooserHandler.MovieRepository,
		chooserHandler.TagRepository)

	output, err := chooserUseCase.AddChoosersToMovieList(dto)
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

func (chooserHandler *WebChooserHandler) AddTagsToMovieList(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputAddTagsToMovieListDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	chooserUseCase := *usecases.NewChooserUseCase(
		chooserHandler.ChooserRepository,
		chooserHandler.MovieListRepository,
		chooserHandler.MovieRepository,
		chooserHandler.TagRepository)

	output, err := chooserUseCase.AddTagsToMovieList(dto)
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
