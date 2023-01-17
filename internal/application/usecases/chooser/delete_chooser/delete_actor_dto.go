package deletechooser

type InputDeleteChooserDto struct {
	ID string `json:"id"`
}

type OutputDeleteChooserDto struct {
	IsDeleted bool `json:"is_deleted"`
}
