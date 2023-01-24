package entity

type ChooserRepositoryInterface interface {
	Create(chooser *Chooser) error
	FindAll() ([]Chooser, error)
	Find(id string) (Chooser, error)
	// Update(c *Chooser) (*Chooser, error)
	// Delete(id string) (*Chooser, error)
}
