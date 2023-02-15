package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebWriterHandler struct {
	WriterRepository entity.WriterRepositoryInterface
	MovieRepository  entity.MovieRepositoryInterface
}

func NewWriterHandler(writerRepository entity.WriterRepositoryInterface, movieRepository entity.MovieRepositoryInterface) *WebWriterHandler {
	return &WebWriterHandler{
		WriterRepository: writerRepository,
		MovieRepository:  movieRepository,
	}
}

func (writerHandler *WebWriterHandler) Create(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputCreateWriterDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository)

	output, err := writerUseCase.Create(dto)
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

func (writerHandler *WebWriterHandler) Find(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	writerId := r.URL.Query().Get("writer_id")

	input := usecases.InputFindWriterDto{
		WriterId: writerId,
	}

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository)

	writer, err := writerUseCase.Find(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(writer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (writerHandler *WebWriterHandler) Update(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPut
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var input usecases.InputUpdateWriterDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository)

	writer, err := writerUseCase.Update(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(writer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (writerHandler *WebWriterHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPatch
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	writerId := r.URL.Query().Get("writer_id")

	input := usecases.InputDeleteWriterDto{
		WriterId: writerId,
	}

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository)

	writer, err := writerUseCase.Delete(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(writer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (writerHandler *WebWriterHandler) IsDeleted(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	writerId := r.URL.Query().Get("writer_id")

	input := usecases.InputIsDeletedWriterDto{
		WriterId: writerId,
	}

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository)

	writer, err := writerUseCase.IsDeleted(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(writer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (writerHandler *WebWriterHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository)

	writers, err := writerUseCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(writers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
