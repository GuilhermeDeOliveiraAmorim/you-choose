package usecases

type MovieDto struct {
	ID              string  `json:"movie_id"`
	Title           string  `json:"title"`
	Synopsis        string  `json:"synopsis"`
	ImdbRating      string  `json:"imdb_rating"`
	Votes           int32   `json:"votes"`
	YouChooseRating float32 `json:"you_choose_rating"`
	Poster          string  `json:"poster"`
	IsDeleted       bool    `json:"is_deleted"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	DeletedAt       string  `json:"deleted_at"`
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

type InputCreateMovieWithImdbIdDto struct {
	ImdbId string `json:"imdb_id"`
}

type OutputCreateMovieWithImdbIdDto struct {
	Movie MovieDto `json:"movie"`
}

type InputFindMovieDto struct {
	ID string `json:"movie_id"`
}

type OutpuFindMovieDto struct {
	ID              string  `json:"movie_id"`
	Title           string  `json:"title"`
	Synopsis        string  `json:"synopsis"`
	ImdbRating      string  `json:"imdb_rating"`
	Votes           int32   `json:"votes"`
	YouChooseRating float32 `json:"you_choose_rating"`
	Poster          string  `json:"poster"`
	IsDeleted       bool    `json:"is_deleted"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	DeletedAt       string  `json:"deleted_at"`
}
