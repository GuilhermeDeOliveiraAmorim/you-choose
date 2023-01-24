package entity

type ChooserRepositoryInterface interface {
	Create(c *Chooser) error
	FindAll() ([]Chooser, error)
	// Update(c *Chooser) (*Chooser, error)
	// Find(id string) (*Chooser, error)
	// Delete(id string) (*Chooser, error)
}
