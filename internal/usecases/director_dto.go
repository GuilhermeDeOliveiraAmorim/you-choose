package usecases

type DirectorDto struct {
	ID        string `json:"director_id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type InputCreateDirectorDto struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputCreateDirectorDto struct {
	Director DirectorDto `json:"director"`
}

type InputDeleteDirectorDto struct {
	ID string `json:"director_id"`
}

type OutputDeleteDirectorDto struct {
	IsDeleted bool `json:"is_deleted"`
}

type InputFindDirectorDto struct {
	ID string `json:"director_id"`
}

type OutputFindDirectorDto struct {
	Director DirectorDto `json:"director"`
}

type InputUpdateDirectorDto struct {
	ID      string `json:"director_id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputUpdateDirectorDto struct {
	Director DirectorDto `json:"director"`
}

type InputIsDeletedDirectorDto struct {
	ID string `json:"director_id"`
}

type OutputIsDeletedDirectorDto struct {
	IsDeleted bool `json:"is_director_deleted"`
}

type OutputFindAllDirectorDto struct {
	Directors []DirectorDto `json:"directors"`
}

type InputFindAllDirectorMoviesDto struct {
	ID string `json:"director_id"`
}

type OutputFindAllDirectorMoviesDto struct {
	Director DirectorDto `json:"director"`
	Movies   []MovieDto  `json:"movies"`
}
