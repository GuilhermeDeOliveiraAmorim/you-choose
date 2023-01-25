package usecases

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
