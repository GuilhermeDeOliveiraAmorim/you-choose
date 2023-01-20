package entity

type DirectorRepositoryInterface interface {
	Create(a *Director) error
	Update(a *Director) error
	Find(id string) (*Director, error)
	Delete(id string) error
	FindAll() ([]*Director, error)
}
