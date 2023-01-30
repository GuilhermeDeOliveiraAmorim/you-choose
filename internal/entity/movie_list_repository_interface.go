package entity

type MovieListRepositoryInterface interface {
	Create(movie *MovieList) error
	FindAll() ([]MovieList, error)
	Find(id string) (MovieList, error)
	AddChooserToMovieList(movieList *MovieList, chooserId *Chooser, created_at string, updated_at string, deleted_at string) error
}
