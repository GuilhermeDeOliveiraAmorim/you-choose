package usecases

type InputCreateChooserDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Picture   string `json:"picture"`
}

type OutputCreateChooserDto struct {
	ID       string `json:"chooser_id"`
	UserName string `json:"username"`
	Picture  string `json:"picture"`
}

type InputDeleteChooserDto struct {
	ID string `json:"chooser_id"`
}

type OutputDeleteChooserDto struct {
	IsDeleted bool `json:"chooser_has_been_deleted"`
}

type InputFindChooserDto struct {
	ID string `json:"chooser_id"`
}

type OutputFindChooserDto struct {
	ID       string `json:"chooser_id"`
	UserName string `json:"username"`
	Picture  string `json:"picture"`
}

type OutputFindAllChooserDto struct {
	Choosers []OutputFindChooserDto `json:"choosers"`
}

type InputUpdateChooserDto struct {
	ID        string `json:"chooser_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Picture   string `json:"picture"`
}

type OutputUpdateChooserDto struct {
	ID        string `json:"chooser_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Picture   string `json:"picture"`
}

type InputIsDeletedChooserDto struct {
	ID string `json:"chooser_id"`
}

type OutputIsDeletedChooserDto struct {
	IsDeleted bool `json:"is_chooser_deleted"`
}

type InputChooser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Picture   string `json:"picture"`
}

type InputMovieList struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
}

type InputCreateChooserAndMovieListDto struct {
	Chooser   InputChooser   `json:"chooser"`
	MovieList InputMovieList `json:"movie_list"`
}

type OutputChooser struct {
	ChooserId string `json:"chooser_id"`
	UserName  string `json:"username"`
	Picture   string `json:"picture"`
}

type OutputMovieList struct {
	ID          string          `json:"movie_list_id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Picture     string          `json:"picture"`
	Choosers    []OutputChooser `json:"choosers"`
}

type OutputCreateChooserAndMovieListDto struct {
	MovieList OutputMovieList `json:"movie_list"`
}

type InputChooserCreateMovieListDto struct {
	ChooserId string         `json:"chooser_id"`
	MovieList InputMovieList `json:"movie_list"`
}

type OutputChooserCreateMovieListDto struct {
	ID          string        `json:"movie_list_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Picture     string        `json:"picture"`
	Chooser     OutputChooser `json:"chooser"`
}

type InputFindAllChooserMovieListsDto struct {
	ChooserId string `json:"chooser_id"`
}

type OutputFindAllChooserMovieListsDto struct {
	Chooser    OutputChooser     `json:"chooser"`
	MovieLists []OutputMovieList `json:"movie_lists"`
}

// {
// 	chooser: {
// 		chooser_id,
// 		username,
// 		picture,
// 		movie_lists: [
// 			{
// 				movie_list_id,
// 				title,
// 				description,
// 				picture,
// 				choosers: [
// 					{
// 						chooser_id,
// 						username,
// 						picture,
// 					},
// 					{
// 						chooser_id,
// 						username,
// 						picture,
// 					},
// 				]
// 			},
// 			{
// 				movie_list_id,
// 				title,
// 				description,
// 				picture,
// 				choosers: [
// 					{
// 						chooser_id,
// 						username,
// 						picture,
// 					},
// 					{
// 						chooser_id,
// 						username,
// 						picture,
// 					},
// 				]
// 			},
// 		]
// 	}
// }
