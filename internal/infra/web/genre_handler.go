package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebGenreHandler struct {
	GenreRepository entity.GenreRepositoryInterface
	MovieRepository entity.MovieRepositoryInterface
	FileRepository  entity.FileRepositoryInterface
}

func NewGenreHandler(
	genreRepository entity.GenreRepositoryInterface,
	movieRepository entity.MovieRepositoryInterface,
	fileRepository  entity.FileRepositoryInterface,
	) *WebGenreHandler {
	return &WebGenreHandler{
		GenreRepository: genreRepository,
		MovieRepository: movieRepository,
		FileRepository: fileRepository,
	}
}

func (genreHandler *WebGenreHandler) Create(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputCreateGenreDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	genreUseCase := *usecases.NewGenreUseCase(genreHandler.GenreRepository, genreHandler.MovieRepository, genreHandler.FileRepository)

	output, err := genreUseCase.Create(dto)
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

func (genreHandler *WebGenreHandler) Find(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	genreId := r.URL.Query().Get("genre_id")

	input := usecases.InputFindGenreDto{
		GenreId: genreId,
	}

	genreUseCase := *usecases.NewGenreUseCase(genreHandler.GenreRepository, genreHandler.MovieRepository, genreHandler.FileRepository)

	genre, err := genreUseCase.Find(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (genreHandler *WebGenreHandler) Update(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPut
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var input usecases.InputUpdateGenreDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	genreUseCase := *usecases.NewGenreUseCase(genreHandler.GenreRepository, genreHandler.MovieRepository, genreHandler.FileRepository)

	genre, err := genreUseCase.Update(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (genreHandler *WebGenreHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPatch
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	genreId := r.URL.Query().Get("genre_id")

	input := usecases.InputDeleteGenreDto{
		GenreId: genreId,
	}

	genreUseCase := *usecases.NewGenreUseCase(genreHandler.GenreRepository, genreHandler.MovieRepository, genreHandler.FileRepository)

	genre, err := genreUseCase.Delete(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (genreHandler *WebGenreHandler) IsDeleted(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	genreId := r.URL.Query().Get("genre_id")

	input := usecases.InputIsDeletedGenreDto{
		GenreId: genreId,
	}

	genreUseCase := *usecases.NewGenreUseCase(genreHandler.GenreRepository, genreHandler.MovieRepository, genreHandler.FileRepository)

	genre, err := genreUseCase.IsDeleted(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (genreHandler *WebGenreHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	genreUseCase := *usecases.NewGenreUseCase(genreHandler.GenreRepository, genreHandler.MovieRepository, genreHandler.FileRepository)

	genres, err := genreUseCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(genres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (genreHandler *WebGenreHandler) AddPictureToGenre(w http.ResponseWriter, r *http.Request) {
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

	var dto usecases.InputAddPictureToGenreDto

	dto.GenreId = r.MultipartForm.Value["genre_id"][0]
	dto.File.File = file
	dto.File.Handler = handler

	genreUseCase := *usecases.NewGenreUseCase(genreHandler.GenreRepository, genreHandler.MovieRepository, genreHandler.FileRepository)

	genres, err := genreUseCase.AddPictureToGenre(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(genres)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (genreHandler *WebGenreHandler) FindGenrePictureToBase64(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	genreId := r.URL.Query().Get("genre_id")

	input := usecases.InputFindGenrePictureToBase64Dto{
		GenreId: genreId,
	}

	genreUseCase := *usecases.NewGenreUseCase(genreHandler.GenreRepository, genreHandler.MovieRepository, genreHandler.FileRepository)

	genre, err := genreUseCase.FindGenrePictureToBase64(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(genre)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
