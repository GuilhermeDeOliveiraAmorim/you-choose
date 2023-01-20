package entity

type WriterRepositoryInterface interface {
	Create(a *Writer) (*Writer, error)
	Update(a *Writer) (*Writer, error)
	Find(id string) (*Writer, error)
	Delete(id string) (*Writer, error)
	FindAll() ([]*Writer, error)
}
