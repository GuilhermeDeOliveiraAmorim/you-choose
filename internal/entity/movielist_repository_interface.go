package entity

type MovieListRepositoryInterface interface {
	Create(movieList *MovieList) error
	Find(id string) (MovieList, error)
	Update(movieList *MovieList) error
	Delete(movieList *MovieList) error
	IsDeleted(id string) error
	FindAll() ([]MovieList, error)
	FindMovieListMovies(id string) ([]string, error)
	FindMovieListChoosers(id string) ([]string, error)
	FindMovieListTags(id string) ([]string, error)
}
