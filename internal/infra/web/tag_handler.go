package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebTagHandler struct {
	TagRepository       entity.TagRepositoryInterface
	MovieListRepository entity.MovieListRepositoryInterface
}

func NewTagHandler(
	tagRepository entity.TagRepositoryInterface,
	movieRepository entity.MovieListRepositoryInterface) *WebTagHandler {
	return &WebTagHandler{
		TagRepository:       tagRepository,
		MovieListRepository: movieRepository,
	}
}

func (tagHandler *WebTagHandler) Create(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputCreateTagDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tagUseCase := *usecases.NewTagUseCase(tagHandler.TagRepository, tagHandler.MovieListRepository)

	output, err := tagUseCase.Create(dto)
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

func (tagHandler *WebTagHandler) Find(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	tagId := r.URL.Query().Get("tag_id")

	input := usecases.InputFindTagDto{
		TagId: tagId,
	}

	tagUseCase := *usecases.NewTagUseCase(tagHandler.TagRepository, tagHandler.MovieListRepository)

	tag, err := tagUseCase.Find(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (tagHandler *WebTagHandler) Update(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPut
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var input usecases.InputUpdateTagDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tagUseCase := *usecases.NewTagUseCase(tagHandler.TagRepository, tagHandler.MovieListRepository)

	tag, err := tagUseCase.Update(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (tagHandler *WebTagHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPatch
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	tagId := r.URL.Query().Get("tag_id")

	input := usecases.InputDeleteTagDto{
		TagId: tagId,
	}

	tagUseCase := *usecases.NewTagUseCase(tagHandler.TagRepository, tagHandler.MovieListRepository)

	tag, err := tagUseCase.Delete(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (tagHandler *WebTagHandler) IsDeleted(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	tagId := r.URL.Query().Get("tag_id")

	input := usecases.InputIsDeletedTagDto{
		TagId: tagId,
	}

	tagUseCase := *usecases.NewTagUseCase(tagHandler.TagRepository, tagHandler.MovieListRepository)

	tag, err := tagUseCase.IsDeleted(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (tagHandler *WebTagHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	tagUseCase := *usecases.NewTagUseCase(tagHandler.TagRepository, tagHandler.MovieListRepository)

	tags, err := tagUseCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
