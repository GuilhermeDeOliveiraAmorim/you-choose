package createmovielist

type InputCreateMovieListDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

type OutputCreateMovieListDto struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}
