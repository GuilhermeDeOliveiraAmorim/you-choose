package entity

type WriterRepositoryInterface interface {
	Create(writer *Writer) error
	Find(id string) (Writer, error)
	Update(writer *Writer) error
	Delete(writer *Writer) error
	IsDeleted(id string) error
	FindAll() ([]Writer, error)
}
