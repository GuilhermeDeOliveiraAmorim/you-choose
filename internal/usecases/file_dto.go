package usecases

import "mime/multipart"

type FileDto struct {
	ID        string `json:"file_id"`
	EntityId  string `json:"entity_id"`
	Name      string `json:"name"`
	Size      string `json:"size"`
	Extension string `json:"extension"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

type InputCreateFileDto struct {
	EntityId string                `json:"entity_id"`
	Name     string                `json:"name"`
	File     multipart.File        `json:"file"`
	Handler  *multipart.FileHeader `json:"handler"`
}

type OutputCreateFileDto struct {
	File FileDto `json:"file"`
}

type InputDeleteFileDto struct {
	FileId string `json:"file_id"`
}

type OutputDeleteFileDto struct {
	IsDeleted bool `json:"is_deleted"`
}

type InputFindFileDto struct {
	FileId string `json:"file_id"`
}

type OutputFindFileDto struct {
	File FileDto `json:"file"`
}

type InputUpdateFileDto struct {
	FileId string `json:"file_id"`
	Name   string `json:"name"`
}

type OutputUpdateFileDto struct {
	File FileDto `json:"file"`
}

type InputIsDeletedFileDto struct {
	FileId string `json:"file_id"`
}

type OutputIsDeletedFileDto struct {
	IsDeleted bool `json:"is_file_deleted"`
}

type OutputFindAllFileDto struct {
	Files []FileDto `json:"files"`
}
