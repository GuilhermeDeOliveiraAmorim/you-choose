package entity

type MovieListRepositoryInterface interface {
	Create(a *MovieList) error
	Update(a *MovieList) error
	Find(id string) (*MovieList, error)
	Delete(id string) error
	FindAll() ([]*MovieList, error)
}
