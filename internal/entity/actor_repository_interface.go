package entity

type ActorRepositoryInterface interface {
	Create(actor *Actor) error
	Find(id string) (Actor, error)
	Update(actor *Actor) error
	Delete(actor *Actor) error
	IsDeleted(id string) error
	FindAll() ([]Actor, error)
	// AddFileToActor(id string) error
}
