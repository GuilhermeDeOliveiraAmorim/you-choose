package deleteactor

type InputDeleteActorDto struct {
	ID string `json:"id"`
}

type OutputDeleteActorDto struct {
	IsDeleted bool `json:"is_deleted"`
}
