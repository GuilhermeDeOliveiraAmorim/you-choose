package entity

type ActorRepositoryInterface interface {
	Create(a *Actor) (*Actor, error)
	Update(a *Actor) (*Actor, error)
	Find(id string) (*Actor, error)
	Delete(id string) (*Actor, error)
	FindAll() ([]*Actor, error)
}
