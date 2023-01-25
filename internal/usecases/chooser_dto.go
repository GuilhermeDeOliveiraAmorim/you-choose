package usecases

type InputCreateChooserDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Picture   string `json:"picture"`
	Password  string `json:"password"`
}

type OutputCreateChooserDto struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Picture  string `json:"picture"`
}

type InputDeleteChooserDto struct {
	ID string `json:"id"`
}

type OutputDeleteChooserDto struct {
	IsDeleted bool `json:"is_deleted"`
}

type InputFindChooserDto struct {
	ID string `json:"id"`
}

type OutputFindChooserDto struct {
	ID       string `json:"id"`
	UserName string `json:"username"`
	Picture  string `json:"picture"`
}

type OutputFindAllChooserDto struct {
	Choosers []OutputFindChooserDto `json:"choosers"`
}
