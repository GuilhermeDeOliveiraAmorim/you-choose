package entity

type GenreRepositoryInterface interface {
	Create(genre *Genre) error
	Find(id string) (Genre, error)
	FindGenreByName(name string) (Genre, error)
	Update(genre *Genre) error
	Delete(genre *Genre) error
	IsDeleted(id string) error
	FindAll() ([]Genre, error)
}
