package usecases

type GenreDto struct {
	ID        string `json:"genre_id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type InputCreateGenreDto struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputCreateGenreDto struct {
	Genre GenreDto `json:"genre"`
}

type InputDeleteGenreDto struct {
	ID string `json:"genre_id"`
}

type OutputDeleteGenreDto struct {
	IsDeleted bool `json:"is_deleted"`
}

type InputFindGenreDto struct {
	ID string `json:"genre_id"`
}

type OutputFindGenreDto struct {
	Genre GenreDto `json:"genre"`
}

type InputUpdateGenreDto struct {
	ID      string `json:"genre_id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputUpdateGenreDto struct {
	Genre GenreDto `json:"genre"`
}

type InputIsDeletedGenreDto struct {
	ID string `json:"genre_id"`
}

type OutputIsDeletedGenreDto struct {
	IsDeleted bool `json:"is_genre_deleted"`
}

type OutputFindAllGenreDto struct {
	Genres []GenreDto `json:"genres"`
}

type InputFindAllGenreMoviesDto struct {
	ID string `json:"genre_id"`
}

type OutputFindAllGenreMoviesDto struct {
	Genre  GenreDto   `json:"genre"`
	Movies []MovieDto `json:"movies"`
}
