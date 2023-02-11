package entity

type ActorRepositoryInterface interface {
	Create(actor *Actor) error                     //
	Update(actor *Actor) error                     //
	Delete(id string) (*Actor, error)              //
	IsDeleted(id string) error                     //
	Find(id string) (Actor, error)                 //
	FindAll() ([]Actor, error)                     //
	FindAllActorMovies(id string) ([]Movie, error) //
}
