package entity

type MovieRepositoryInterface interface {
	Create(movie *Movie) error
	FindAll() ([]Movie, error)
	Find(id string) (Movie, error)
	AddActorsToMovie(movie Movie, actors []Actor) error
	FindMovieActors(id string) ([]string, error)
	AddWritersToMovie(movie Movie, writers []Writer) error
	FindMovieWriters(id string) ([]string, error)
}
