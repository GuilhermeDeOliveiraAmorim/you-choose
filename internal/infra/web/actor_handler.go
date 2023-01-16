package web

import (
	"encoding/json"
	"net/http"

	actorCreate "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/application/usecases/actor/create_actor"
	actor "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/repository"
)

type WebActorHandler struct {
	ActorRepository actor.ActorRepositoryInterface
}

func NewWebActorHandler(ActorRepository actor.ActorRepositoryInterface) *WebActorHandler {
	return &WebActorHandler{
		ActorRepository: ActorRepository,
	}
}

func (h *WebActorHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto actorCreate.InputCreateActorDto

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createActor := actorCreate.NewCreateActorUseCase(h.ActorRepository)
	output, err := createActor.Execute(dto)

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
