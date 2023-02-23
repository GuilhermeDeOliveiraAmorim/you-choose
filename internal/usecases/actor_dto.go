package usecases

type ActorDto struct {
	ID        string `json:"actor_id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type InputCreateActorDto struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputCreateActorDto struct {
	Actor ActorDto `json:"actor"`
}

type InputDeleteActorDto struct {
	ActorId string `json:"actor_id"`
}

type OutputDeleteActorDto struct {
	IsDeleted bool `json:"is_deleted"`
}

type InputFindActorDto struct {
	ActorId string `json:"actor_id"`
}

type OutputFindActorDto struct {
	Actor ActorDto `json:"actor"`
}

type InputUpdateActorDto struct {
	ActorId string `json:"actor_id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputUpdateActorDto struct {
	Actor ActorDto `json:"actor"`
}

type InputIsDeletedActorDto struct {
	ActorId string `json:"actor_id"`
}

type OutputIsDeletedActorDto struct {
	IsDeleted bool `json:"is_actor_deleted"`
}

type OutputFindAllActorDto struct {
	Actors []ActorDto `json:"actors"`
}

type InputFindAllActorMoviesDto struct {
	ActorId string `json:"actor_id"`
}

type OutputFindAllActorMoviesDto struct {
	Actor  ActorDto   `json:"actor"`
	Movies []MovieDto `json:"movies"`
}

type InputAddFileToActorDto struct {
}

type OutputAddFileToActorDto struct {
}
