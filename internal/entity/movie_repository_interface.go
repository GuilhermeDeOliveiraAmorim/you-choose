package entity

type MovieRepositoryInterface interface {
	UpdateYouChooseRating(id string) (*Movie, error)
	Create(a *Movie) (*Movie, error)
	Update(a *Movie) (*Movie, error)
	Delete(id string) (*Movie, error)
	Find(id string) (*Movie, error)
	FindAll() ([]*Movie, error)
}
