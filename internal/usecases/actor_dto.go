package usecases

type ActorDto struct {
	ID        string  `json:"actor_id"`
	Name      string  `json:"name"`
	Picture   string  `json:"picture"`
	IsDeleted bool    `json:"is_deleted"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
	File      FileDto `json:"file"`
}

type InputCreateActorDto struct {
	Name    string `json:"name"`
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

type InputAddPictureToActorDto struct {
	ActorId string             `json:"actor_id"`
	File    InputCreateFileDto `json:"file"`
}

type OutputAddPictureToActorDto struct {
	Actor ActorDto `json:"actor"`
}

type InputFindActorPictureToBase64Dto struct {
	ActorId string `json:"actor_id"`
}

type OutputFindActorPictureToBase64Dto struct {
	Actor ActorDto `json:"actor"`
}
