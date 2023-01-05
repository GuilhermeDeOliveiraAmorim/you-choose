package createchooser

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
