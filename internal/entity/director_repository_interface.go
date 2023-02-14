package entity

type DirectorRepositoryInterface interface {
	Create(director *Director) error
	Find(id string) (Director, error)
	Update(director *Director) error
	Delete(director *Director) error
	IsDeleted(id string) error
	FindAll() ([]Director, error)
}
