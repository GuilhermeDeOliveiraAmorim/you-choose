package entity

type ChooserRepositoryInterface interface {
	Create(chooser *Chooser) error
	Find(id string) (Chooser, error)
	Update(chooser *Chooser) error
	Delete(chooser *Chooser) error
	IsDeleted(id string) error
	FindAll() ([]Chooser, error)
}
