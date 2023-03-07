package usecases

type WriterDto struct {
	ID        string `json:"writer_id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	File      FileDto `json:"file"`
}

type InputCreateWriterDto struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputCreateWriterDto struct {
	Writer WriterDto `json:"writer"`
}

type InputDeleteWriterDto struct {
	WriterId string `json:"writer_id"`
}

type OutputDeleteWriterDto struct {
	IsDeleted bool `json:"is_deleted"`
}

type InputFindWriterDto struct {
	WriterId string `json:"writer_id"`
}

type OutputFindWriterDto struct {
	Writer WriterDto `json:"writer"`
}

type InputUpdateWriterDto struct {
	WriterId string `json:"writer_id"`
	Name     string `json:"name"`
	Picture  string `json:"picture"`
}

type OutputUpdateWriterDto struct {
	Writer WriterDto `json:"writer"`
}

type InputIsDeletedWriterDto struct {
	WriterId string `json:"writer_id"`
}

type OutputIsDeletedWriterDto struct {
	IsDeleted bool `json:"is_writer_deleted"`
}

type OutputFindAllWriterDto struct {
	Writers []WriterDto `json:"writers"`
}

type InputFindAllWriterMoviesDto struct {
	WriterId string `json:"writer_id"`
}

type OutputFindAllWriterMoviesDto struct {
	Writer WriterDto  `json:"writer"`
	Movies []MovieDto `json:"movies"`
}


type InputAddPictureToWriterDto struct {
	WriterId string             `json:"writer_id"`
	File    InputCreateFileDto `json:"file"`
}

type OutputAddPictureToWriterDto struct {
	Writer WriterDto `json:"writer"`
}

type InputFindWriterPictureToBase64Dto struct {
	WriterId string `json:"writer_id"`
}

type OutputFindWriterPictureToBase64Dto struct {
	Writer WriterDto `json:"writer"`
}
