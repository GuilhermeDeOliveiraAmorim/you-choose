package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebDirectorHandler struct {
	DirectorRepository entity.DirectorRepositoryInterface
	MovieRepository    entity.MovieRepositoryInterface
	FileRepository  entity.FileRepositoryInterface
}

func NewDirectorHandler(
	directorRepository entity.DirectorRepositoryInterface,
	movieRepository entity.MovieRepositoryInterface,
	fileRepository entity.FileRepositoryInterface) *WebDirectorHandler {
	return &WebDirectorHandler{
		DirectorRepository: directorRepository,
		MovieRepository:    movieRepository,
		FileRepository:  fileRepository,
	}
}

func (directorHandler *WebDirectorHandler) Create(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputCreateDirectorDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	directorUseCase := *usecases.NewDirectorUseCase(directorHandler.DirectorRepository, directorHandler.MovieRepository, directorHandler.FileRepository)

	output, err := directorUseCase.Create(dto)
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

func (directorHandler *WebDirectorHandler) Find(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	directorId := r.URL.Query().Get("director_id")

	input := usecases.InputFindDirectorDto{
		DirectorId: directorId,
	}

	directorUseCase := *usecases.NewDirectorUseCase(directorHandler.DirectorRepository, directorHandler.MovieRepository, directorHandler.FileRepository)

	director, err := directorUseCase.Find(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(director)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (directorHandler *WebDirectorHandler) Update(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPut
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var input usecases.InputUpdateDirectorDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	directorUseCase := *usecases.NewDirectorUseCase(directorHandler.DirectorRepository, directorHandler.MovieRepository, directorHandler.FileRepository)

	director, err := directorUseCase.Update(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(director)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (directorHandler *WebDirectorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPatch
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	directorId := r.URL.Query().Get("director_id")

	input := usecases.InputDeleteDirectorDto{
		DirectorId: directorId,
	}

	directorUseCase := *usecases.NewDirectorUseCase(directorHandler.DirectorRepository, directorHandler.MovieRepository, directorHandler.FileRepository)

	director, err := directorUseCase.Delete(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(director)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (directorHandler *WebDirectorHandler) IsDeleted(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	directorId := r.URL.Query().Get("director_id")

	input := usecases.InputIsDeletedDirectorDto{
		DirectorId: directorId,
	}

	directorUseCase := *usecases.NewDirectorUseCase(directorHandler.DirectorRepository, directorHandler.MovieRepository, directorHandler.FileRepository)

	director, err := directorUseCase.IsDeleted(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(director)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (directorHandler *WebDirectorHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	directorUseCase := *usecases.NewDirectorUseCase(directorHandler.DirectorRepository, directorHandler.MovieRepository, directorHandler.FileRepository)

	directors, err := directorUseCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(directors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
