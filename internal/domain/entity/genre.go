package entity

type Genre struct {
	ID      string
	Name    string
	Picture string
}

func NewGenre(id, name string, picture string) *Genre {
	return &Genre{
		ID:      id,
		Name:    name,
		Picture: picture,
	}
}
