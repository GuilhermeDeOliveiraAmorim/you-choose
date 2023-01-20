package entity

type GenreRepositoryInterface interface {
	Create(a *Genre) (*Genre, error)
	Update(a *Genre) (*Genre, error)
	Find(id string) (*Genre, error)
	Delete(id string) (*Genre, error)
	FindAll() ([]*Genre, error)
}
