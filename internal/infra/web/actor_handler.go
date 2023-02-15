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
}

func NewActorHandler(actorRepository entity.ActorRepositoryInterface, movieRepository entity.MovieRepositoryInterface) *WebActorHandler {
	return &WebActorHandler{
		ActorRepository: actorRepository,
		MovieRepository: movieRepository,
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

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository)

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

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository)

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

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository)

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

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository)

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

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository)

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

	actorUseCase := *usecases.NewActorUseCase(actorHandler.ActorRepository, actorHandler.MovieRepository)

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
