package entity

type ChooserRepositoryInterface interface {
	Create(chooser *Chooser) error
	Update(chooser *Chooser) error
	Delete(chooser *Chooser) error
	Find(id string) (Chooser, error)
	FindAll() ([]Chooser, error)
	IsDeleted(id string) error
	CreateChooserAndMovieList(chooser *Chooser, movieList *MovieList) error
	ChooserCreateMovieList(chooser *Chooser, movieList *MovieList) error
	FindAllChooserMovieLists(id string) ([]MovieList, error)
}