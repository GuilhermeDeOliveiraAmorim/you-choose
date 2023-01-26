package entity

type MovieListRepositoryInterface interface {
	Create(movie *MovieList) error
	FindAll() ([]MovieList, error)
	Find(id string) (MovieList, error)
	// Update(a *MovieList) error
	// Delete(id string) error
}
