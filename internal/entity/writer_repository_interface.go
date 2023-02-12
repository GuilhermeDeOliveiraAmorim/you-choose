package entity

type WriterRepositoryInterface interface {
	Create(writer *Writer) error
	Find(id string) (Writer, error)
	Update(writer *Writer) error
	Delete(id string) (*Writer, error)
	IsDeleted(id string) error
	FindAll() ([]Writer, error)
	FindAllWriterMovies(id string) ([]Movie, error)
}
