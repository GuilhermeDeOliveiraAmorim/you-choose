package createactor

type InputCreateActorDto struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputCreateActorDto struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
