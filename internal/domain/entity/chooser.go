package entity

type Chooser struct {
	ID        string
	FirstName string
	LastName  string
	UserName  string
	Picture   string
}

func NewChooser(id string, firstName string, lastName string, userName string, picture string) *Chooser {
	return &Chooser{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
		Picture:   picture,
	}
}
