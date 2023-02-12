package usecases

type WriterDto struct {
	ID        string `json:"writer_id"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type InputCreateWriterDto struct {
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputCreateWriterDto struct {
	Writer WriterDto `json:"writer"`
}

type InputDeleteWriterDto struct {
	ID string `json:"writer_id"`
}

type OutputDeleteWriterDto struct {
	IsDeleted bool `json:"is_deleted"`
}

type InputFindWriterDto struct {
	ID string `json:"writer_id"`
}

type OutputFindWriterDto struct {
	Writer WriterDto `json:"writer"`
}

type InputUpdateWriterDto struct {
	ID      string `json:"writer_id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OutputUpdateWriterDto struct {
	Writer WriterDto `json:"writer"`
}

type InputIsDeletedWriterDto struct {
	ID string `json:"writer_id"`
}

type OutputIsDeletedWriterDto struct {
	IsDeleted bool `json:"is_writer_deleted"`
}

type OutputFindAllWriterDto struct {
	Writers []WriterDto `json:"writers"`
}

type InputFindAllWriterMoviesDto struct {
	ID string `json:"writer_id"`
}

type OutputFindAllWriterMoviesDto struct {
	Writer WriterDto  `json:"writer"`
	Movies []MovieDto `json:"movies"`
}
