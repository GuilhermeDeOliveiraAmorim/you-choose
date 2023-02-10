package usecases

type MovieListDto struct {
	ID              string       `json:"movie_list_id"`
	Title           string       `json:"title"`
	Synopsis        string       `json:"synopsis"`
	ImdbRating      string       `json:"imdb_rating"`
	Votes           int          `json:"votes"`
	YouChooseRating int          `json:"you_choose_rating"`
	Poster          string       `json:"poster"`
	IsDeleted       bool         `json:"is_deleted"`
	CreatedAt       string       `json:"created_at"`
	UpdatedAt       string       `json:"updated_at"`
	DeletedAt       string       `json:"deleted_at"`
	Choosers        []ChooserDto `json:"choosers"`
	Movies          []MovieDto   `json:"movies"`
}

type InputAddChooserToMovieListDto struct {
	ChooserId   string `json:"chooser_id"`
	MovieListId string `json:"movie_list_id"`
}

type OutputAddChooserToMovieListDto struct {
	Chooser   OutputFindChooserDto   `json:"chooser"`
	MovieList OutputFindMovieListDto `json:"movie_list"`
}

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

type InputFindMovieListDto struct {
	ID string `json:"id"`
}

type OutputFindMovieListDto struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

type OutputFindAllMovieListDto struct {
	MovieLists []OutputFindMovieListDto `json:"movie_lists"`
}

type InputFindChooserInMovieList struct {
	ChooserId   string `json:"chooser_id"`
	MovieListId string `json:"movie_list_id"`
}

type OutpuFindChooserInMovieList struct {
	ChooserIsInTheMovieList bool `json:"chooser_is_in_the_movie_list"`
}
