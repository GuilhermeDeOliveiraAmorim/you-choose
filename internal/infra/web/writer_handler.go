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
	FileRepository  entity.FileRepositoryInterface
}

func NewWriterHandler(
	writerRepository entity.WriterRepositoryInterface,
	movieRepository entity.MovieRepositoryInterface,
	fileRepository  entity.FileRepositoryInterface) *WebWriterHandler {
	return &WebWriterHandler{
		WriterRepository: writerRepository,
		MovieRepository:  movieRepository,
		FileRepository: fileRepository,
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

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository, writerHandler.FileRepository)

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

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository, writerHandler.FileRepository)

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

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository, writerHandler.FileRepository)

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

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository, writerHandler.FileRepository)

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

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository, writerHandler.FileRepository)

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

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository, writerHandler.FileRepository)

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


func (writerHandler *WebWriterHandler) AddPictureToWriter(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	err := r.ParseMultipartForm(1 << 2)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var dto usecases.InputAddPictureToWriterDto

	dto.WriterId = r.MultipartForm.Value["writer_id"][0]
	dto.File.File = file
	dto.File.Handler = handler

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository, writerHandler.FileRepository)

	writers, err := writerUseCase.AddPictureToWriter(dto)
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

func (writerHandler *WebWriterHandler) FindWriterPictureToBase64(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	writerId := r.URL.Query().Get("writer_id")

	input := usecases.InputFindWriterPictureToBase64Dto{
		WriterId: writerId,
	}

	writerUseCase := *usecases.NewWriterUseCase(writerHandler.WriterRepository, writerHandler.MovieRepository, writerHandler.FileRepository)

	writer, err := writerUseCase.FindWriterPictureToBase64(input)
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
