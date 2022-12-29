package entity

type Director struct {
	ID      string
	Name    string
	Picture string
}

func NewDirector(id, name string, picture string) *Director {
	return &Director{
		ID:      id,
		Name:    name,
		Picture: picture,
	}
}
