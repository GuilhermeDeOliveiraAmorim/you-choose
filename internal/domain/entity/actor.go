package entity

type Actor struct {
	ID      string
	Name    string
	Picture string
}

func NewActor(id, name string, picture string) *Actor {
	return &Actor{
		ID:      id,
		Name:    name,
		Picture: picture,
	}
}
