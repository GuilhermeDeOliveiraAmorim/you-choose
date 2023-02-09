package entity

type MovieRepositoryInterface interface {
	Create(movie *Movie) error
	FindAll() ([]Movie, error)
	Find(id string) (Movie, error)
}
