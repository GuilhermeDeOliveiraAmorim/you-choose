package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WebActorHandler struct {
	ActorRepository entity.ActorRepositoryInterface
	MovieRepository entity.MovieRepositoryInterface
	FileRepository  entity.FileRepositoryInterface
}

func NewActorHandler(
	actorRepository entity.ActorRepositoryInterface,
	movieRepository entity.MovieRepositoryInterface,
	fileRepository entity.FileRepositoryInterface) *WebActorHandler {
	return &WebActorHandler{
		ActorRepository: actorRepository,
		MovieRepository: movieRepository,
		FileRepository:  fileRepository,
	}
}

func (actorHandler *WebActorHandler) Create(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPost
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var dto usecases.InputCreateActorDto

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository, actorHandler.FileRepository)

	output, err := actorUseCase.Create(dto)
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

func (actorHandler *WebActorHandler) Find(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	actorId := r.URL.Query().Get("actor_id")

	input := usecases.InputFindActorDto{
		ActorId: actorId,
	}

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository, actorHandler.FileRepository)

	actor, err := actorUseCase.Find(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (actorHandler *WebActorHandler) Update(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPut
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	var input usecases.InputUpdateActorDto

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository, actorHandler.FileRepository)

	actor, err := actorUseCase.Update(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (actorHandler *WebActorHandler) Delete(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodPatch
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	actorId := r.URL.Query().Get("actor_id")

	input := usecases.InputDeleteActorDto{
		ActorId: actorId,
	}

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository, actorHandler.FileRepository)

	actor, err := actorUseCase.Delete(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (actorHandler *WebActorHandler) IsDeleted(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	actorId := r.URL.Query().Get("actor_id")

	input := usecases.InputIsDeletedActorDto{
		ActorId: actorId,
	}

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository, actorHandler.FileRepository)

	actor, err := actorUseCase.IsDeleted(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (actorHandler *WebActorHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository, actorHandler.FileRepository)

	actors, err := actorUseCase.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(actors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (actorHandler *WebActorHandler) AddPictureToActor(w http.ResponseWriter, r *http.Request) {
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

	var dto usecases.InputAddPictureToActorDto

	dto.ActorId = r.MultipartForm.Value["actor_id"][0]
	dto.File.File = file
	dto.File.Handler = handler

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository, actorHandler.FileRepository)

	actors, err := actorUseCase.AddPictureToActor(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(actors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (actorHandler *WebActorHandler) FindActorPictureToBase64(w http.ResponseWriter, r *http.Request) {
	handlerMethod := http.MethodGet
	requestMethod := r.Method
	if handlerMethod != requestMethod {
		http.Error(w, requestMethod+" method not allowed", http.StatusInternalServerError)
		return
	}

	actorId := r.URL.Query().Get("actor_id")

	input := usecases.InputFindActorPictureToBase64Dto{
		ActorId: actorId,
	}

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository, actorHandler.FileRepository)

	actor, err := actorUseCase.FindActorPictureToBase64(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

