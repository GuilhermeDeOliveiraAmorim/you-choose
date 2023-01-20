package web

import (
	"encoding/json"
	"net/http"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
)

type WebActorHandler struct {
	ActorRepository entity.ActorRepositoryInterface
}

func NewWebActorHandler(ActorRepository entity.ActorRepositoryInterface) *WebActorHandler {
	return &WebActorHandler{
		ActorRepository: ActorRepository,
	}
}

func (h *WebActorHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecases.InputCreateActorDto

	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createActor := usecases.NewActorUseCase(h.ActorRepository)
	output, err := createActor.Create(dto)

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
