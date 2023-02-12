package entity

type ActorRepositoryInterface interface {
	Create(actor *Actor) error
	Find(id string) (Actor, error)
	Update(actor *Actor) error
	Delete(id string) (*Actor, error)
	IsDeleted(id string) error
	FindAll() ([]Actor, error)
	FindAllActorMovies(id string) ([]Movie, error)
}
