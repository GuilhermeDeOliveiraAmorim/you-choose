package usecases

type InputCreateDirectorDto struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputCreateDirectorDto struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}
