package entity

type Writer struct {
	ID      string
	Name    string
	Picture string
}

func NewWriter(id, name string, picture string) *Writer {
	return &Writer{
		ID:      id,
		Name:    name,
		Picture: picture,
	}
}
