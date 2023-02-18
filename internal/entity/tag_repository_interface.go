package entity

type TagRepositoryInterface interface {
	Create(tag *Tag) error
	Find(id string) (Tag, error)
	FindTagByName(name string) (Tag, error)
	Update(tag *Tag) error
	Delete(tag *Tag) error
	IsDeleted(id string) error
	FindAll() ([]Tag, error)
}
