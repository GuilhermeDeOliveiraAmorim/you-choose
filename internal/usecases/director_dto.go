package usecases

type DirectorDto struct {
	ID        string `json:"director_id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string  `json:"deleted_at"`
	File      FileDto `json:"file"`
}

type InputCreateDirectorDto struct {
	Name    string `json:"name"`
}

type OutputCreateDirectorDto struct {
	Director DirectorDto `json:"director"`
}

type InputDeleteDirectorDto struct {
	DirectorId string `json:"director_id"`
}

type OutputDeleteDirectorDto struct {
	IsDeleted bool `json:"is_deleted"`
}

type InputFindDirectorDto struct {
	DirectorId string `json:"director_id"`
}

type OutputFindDirectorDto struct {
	Director DirectorDto `json:"director"`
}

type InputUpdateDirectorDto struct {
	DirectorId string `json:"director_id"`
	Name       string `json:"name"`
	Picture    string `json:"picture"`
}

type OutputUpdateDirectorDto struct {
	Director DirectorDto `json:"director"`
}

type InputIsDeletedDirectorDto struct {
	DirectorId string `json:"director_id"`
}

type OutputIsDeletedDirectorDto struct {
	IsDeleted bool `json:"is_director_deleted"`
}

type OutputFindAllDirectorDto struct {
	Directors []DirectorDto `json:"directors"`
}

type InputFindAllDirectorMoviesDto struct {
	DirectorId string `json:"director_id"`
}

type OutputFindAllDirectorMoviesDto struct {
	Director DirectorDto `json:"director"`
	Movies   []MovieDto  `json:"movies"`
}

type InputAddPictureToDirectorDto struct {
	DirectorId string             `json:"director_id"`
	File    InputCreateFileDto `json:"file"`
}

type OutputAddPictureToDirectorDto struct {
	Director DirectorDto `json:"director"`
}

type InputFindDirectorPictureToBase64Dto struct {
	DirectorId string `json:"director_id"`
}

type OutputFindDirectorPictureToBase64Dto struct {
	Director DirectorDto `json:"director"`
}
