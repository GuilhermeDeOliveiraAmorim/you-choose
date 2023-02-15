package usecases

type MovieDto struct {
	ID              string        `json:"movie_id"`
	Title           string        `json:"title"`
	Synopsis        string        `json:"synopsis"`
	ImdbRating      string        `json:"imdb_rating"`
	Votes           int32         `json:"votes"`
	YouChooseRating float32       `json:"you_choose_rating"`
	Poster          string        `json:"poster"`
	IsDeleted       bool          `json:"is_deleted"`
	CreatedAt       string        `json:"created_at"`
	UpdatedAt       string        `json:"updated_at"`
	DeletedAt       string        `json:"deleted_at"`
	Actors          []ActorDto    `json:"actors"`
	Writers         []WriterDto   `json:"writers"`
	Directors       []DirectorDto `json:"directors"`
	Genres          []GenreDto    `json:"genres"`
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

type OutputCreateMovieWithImdbIdDto struct {
	Movie MovieDto `json:"movie"`
}

type InputFindMovieDto struct {
	MovieId string `json:"movie_id"`
}

type OutpuFindMovieDto struct {
	Movie MovieDto `json:"movie"`
}

type ActorId struct {
	ActorId string `json:"actor_id"`
}

type InputAddActorsToMovieDto struct {
	MovieId   string    `json:"movie_id"`
	ActorsIds []ActorId `json:"actors_ids"`
}

type OutputAddActorsToMovieDto struct {
	Movie MovieDto `json:"movie"`
}

type InputFindMovieActorsDto struct {
	MovieId string `json:"movie_id"`
}

type OutputFindMovieActorsDto struct {
	Movie MovieDto `json:"movie"`
}

type WriterId struct {
	WriterId string `json:"writer_id"`
}

type InputAddWritersToMovieDto struct {
	MovieId    string     `json:"movie_id"`
	WritersIds []WriterId `json:"writers_ids"`
}

type OutputAddWritersToMovieDto struct {
	Movie MovieDto `json:"movie"`
}

type InputFindMovieWritersDto struct {
	MovieId string `json:"movie_id"`
}

type OutputFindMovieWritersDto struct {
	Movie MovieDto `json:"movie"`
}

type DirectorId struct {
	DirectorId string `json:"director_id"`
}

type InputAddDirectorsToMovieDto struct {
	MovieId      string       `json:"movie_id"`
	DirectorsIds []DirectorId `json:"directors_ids"`
}

type OutputAddDirectorsToMovieDto struct {
	Movie MovieDto `json:"movie"`
}

type InputFindMovieDirectorsDto struct {
	MovieId string `json:"movie_id"`
}

type OutputFindMovieDirectorsDto struct {
	Movie MovieDto `json:"movie"`
}

type GenreId struct {
	GenreId string `json:"genre_id"`
}

type InputAddGenresToMovieDto struct {
	MovieId   string    `json:"movie_id"`
	GenresIds []GenreId `json:"genres_ids"`
}

type OutputAddGenresToMovieDto struct {
	Movie MovieDto `json:"movie"`
}

type InputFindMovieGenresDto struct {
	MovieId string `json:"movie_id"`
}

type OutputFindMovieGenresDto struct {
	Movie MovieDto `json:"movie"`
}
