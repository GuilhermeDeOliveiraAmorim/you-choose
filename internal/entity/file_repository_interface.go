package entity

type FileRepositoryInterface interface {
	Create(file *File) error
	Find(id string) (File, error)
	// Update(file *File) error
	// Delete(file *File) error
	// IsDeleted(id string) error
	// FindAll() ([]File, error)
}
