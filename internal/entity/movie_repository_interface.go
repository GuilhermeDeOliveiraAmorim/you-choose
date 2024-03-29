package entity

type MovieRepositoryInterface interface {
	Create(movie *Movie) error
	Find(id string) (Movie, error)
	Update(movie *Movie) error
	Delete(movie *Movie) error
	IsDeleted(id string) error
	FindAll() ([]Movie, error)
	AddActorsToMovie(movie Movie, actors []Actor) error
	FindMovieActors(id string) ([]string, error)
	AddWritersToMovie(movie Movie, writers []Writer) error
	FindMovieWriters(id string) ([]string, error)
	AddDirectorsToMovie(movie Movie, directors []Director) error
	FindMovieDirectors(id string) ([]string, error)
	AddGenresToMovie(movie Movie, genres []Genre) error
	FindMovieGenres(id string) ([]string, error)
	AddVoteToMovie(movie *Movie) error
}
