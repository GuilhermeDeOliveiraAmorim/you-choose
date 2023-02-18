package usecases

type MovieListDto struct {
	ID          string       `json:"movie_list_id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Picture     string       `json:"picture"`
	IsDeleted   bool         `json:"is_deleted"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
	DeletedAt   string       `json:"deleted_at"`
	Choosers    []ChooserDto `json:"choosers"`
	Movies      []MovieDto   `json:"movies"`
	Tags        []TagDto     `json:"tags"`
}

type InputCreateMovieListDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

type OutputCreateMovieListDto struct {
	MovieList MovieListDto `json:"movie_list"`
}

type InputFindMovieListDto struct {
	MovieListId string `json:"movie_list_id"`
}

type OutputFindMovieListDto struct {
	MovieList MovieListDto `json:"movie_list"`
}

type InputDeleteMovieListDto struct {
	MovieListId string `json:"movie_list_id"`
}

type OutputDeleteMovieListDto struct {
	IsDeleted bool `json:"is_deleted"`
}

type InputUpdateMovieListDto struct {
	MovieListId string `json:"movie_list_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

type OutputUpdateMovieListDto struct {
	MovieList MovieListDto `json:"movie_list"`
}

type InputIsDeletedMovieListDto struct {
	MovieListId string `json:"movie_list_id"`
}

type OutputIsDeletedMovieListDto struct {
	IsDeleted bool `json:"is_movie_list_deleted"`
}

type OutputFindAllMovieListDto struct {
	MovieLists []MovieListDto `json:"movie_lists"`
}

type InputFindMovieListMoviesDto struct {
	MovieListId string `json:"movie_list_id"`
}

type OutputFindMovieListMoviesDto struct {
	MovieList MovieListDto `json:"movie_list"`
}

type InputFindMovieListChoosersDto struct {
	MovieListId string `json:"movie_list_id"`
}

type OutputFindMovieListChoosersDto struct {
	MovieList MovieListDto `json:"movie_list"`
}

type InputFindMovieListTagsDto struct {
	MovieListId string `json:"movie_list_id"`
}

type OutputFindMovieListTagsDto struct {
	MovieList MovieListDto `json:"movie_list"`
}
