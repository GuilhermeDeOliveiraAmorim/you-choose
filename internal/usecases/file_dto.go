package usecases

type FileDto struct {
	ID        string `json:"file_id"`
	Name      string `json:"name"`
	Size      string `json:"size"`
	Extension string `json:"extension"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type InputCreateFileDto struct {
	Name      string `json:"name"`
	Size      string `json:"size"`
	Extension string `json:"extension"`
}

type OutputCreateFileDto struct {
	File FileDto `json:"file"`
}
