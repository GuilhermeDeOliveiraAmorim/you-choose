package usecases

type MovieDto struct {
	ID              string  `json:"movie_id"`
	Title           string  `json:"title"`
	Synopsis        string  `json:"synopsis"`
	ImdbRating      string  `json:"imdb_rating"`
	Votes           int32   `json:"votes"`
	YouChooseRating float32 `json:"you_choose_rating"`
	Poster          string  `json:"poster"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	DeletedAt       string  `json:"deleted_at"`
	IsDeleted       bool    `json:"is_deleted"`
}

type InputCreateMovieDto struct {
	Title      string `json:"title"`
	Synopsis   string `json:"synopsis"`
	ImdbRating string `json:"imdb_rating"`
	Poster     string `json:"poster"`
}

type OutputCreateMovieDto struct {
	Movie MovieDto `json:"movie"`
}

type OutputFindAllMoviesDto struct {
	Movies []MovieDto `json:"movies"`
}