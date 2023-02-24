package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebFileHandler struct {
	FileRepository entity.FileRepositoryInterface
}

func NewFileHandler(fileRepository entity.FileRepositoryInterface) *WebFileHandler {
	return &WebFileHandler{
		FileRepository: fileRepository,
	}
}

func (fileHandler *WebFileHandler) Create(w http.ResponseWriter, r *http.Request) {
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

	var dto usecases.InputCreateFileDto

	dto.EntityId = r.MultipartForm.Value["entity_id"][0]
	dto.Name = r.MultipartForm.Value["name"][0]
	dto.File = file
	dto.Handler = handler

	errs := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, errs.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(dto.EntityId)
	fmt.Println(dto.Name)

	fileUseCase := *usecases.NewFileUseCase(fileHandler.FileRepository)

	output, err := fileUseCase.Create(dto)
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

// func (fileHandler *WebFileHandler) Find(w http.ResponseWriter, r *http.Request) {
// 	handlerMethod := http.MethodGet
// 	requestMethod := r.Method
// 	if handlerMethod != requestMethod {
// 		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
// 		return
// 	}

// 	fileId := r.URL.Query().Get("file_id")

// 	input := usecases.InputFindFileDto{
// 		FileId: fileId,
// 	}

// 	fileUseCase := *usecases.NewFileUseCase(fileHandler.FileRepository)

// 	file, err := fileUseCase.Find(input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(file)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func (fileHandler *WebFileHandler) Update(w http.ResponseWriter, r *http.Request) {
// 	handlerMethod := http.MethodPut
// 	requestMethod := r.Method
// 	if handlerMethod != requestMethod {
// 		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
// 		return
// 	}

// 	var input usecases.InputUpdateFileDto

// 	err := json.NewDecoder(r.Body).Decode(&input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	fileUseCase := *usecases.NewFileUseCase(fileHandler.FileRepository)

// 	file, err := fileUseCase.Update(input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(file)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func (fileHandler *WebFileHandler) Delete(w http.ResponseWriter, r *http.Request) {
// 	handlerMethod := http.MethodPatch
// 	requestMethod := r.Method
// 	if handlerMethod != requestMethod {
// 		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
// 		return
// 	}

// 	fileId := r.URL.Query().Get("file_id")

// 	input := usecases.InputDeleteFileDto{
// 		FileId: fileId,
// 	}

// 	fileUseCase := *usecases.NewFileUseCase(fileHandler.FileRepository)

// 	file, err := fileUseCase.Delete(input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(file)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func (fileHandler *WebFileHandler) IsDeleted(w http.ResponseWriter, r *http.Request) {
// 	handlerMethod := http.MethodGet
// 	requestMethod := r.Method
// 	if handlerMethod != requestMethod {
// 		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
// 		return
// 	}

// 	fileId := r.URL.Query().Get("file_id")

// 	input := usecases.InputIsDeletedFileDto{
// 		FileId: fileId,
// 	}

// 	fileUseCase := *usecases.NewFileUseCase(fileHandler.FileRepository)

// 	file, err := fileUseCase.IsDeleted(input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(file)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func (fileHandler *WebFileHandler) FindAll(w http.ResponseWriter, r *http.Request) {
// 	handlerMethod := http.MethodGet
// 	requestMethod := r.Method
// 	if handlerMethod != requestMethod {
// 		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
// 		return
// 	}

// 	fileUseCase := *usecases.NewFileUseCase(fileHandler.FileRepository)

// 	files, err := fileUseCase.FindAll()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	err = json.NewEncoder(w).Encode(files)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
