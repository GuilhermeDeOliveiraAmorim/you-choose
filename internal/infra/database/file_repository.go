package database

import (
	"database/sql"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type FileRepository struct {
	Db *sql.DB
}

func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{
		Db: db,
	}
}

func (fileRepository *FileRepository) Create(file *entity.File) error {
	stmt, err := fileRepository.Db.Prepare("INSERT INTO files (file_id, entity_id, name, size, extension, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(file.ID, file.EntityId, file.Name, file.Size, file.Extension, file.IsDeleted, file.CreatedAt, file.UpdatedAt, file.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (fileRepository *FileRepository) Find(id string) (entity.File, error) {
	var file entity.File

	rows, err := fileRepository.Db.Query("SELECT * FROM files WHERE file_id = $1", id)
	if err != nil {
		return file, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&file.ID,
			&file.EntityId,
			&file.Name,
			&file.Size,
			&file.Extension,
			&file.IsDeleted,
			&file.CreatedAt,
			&file.UpdatedAt,
			&file.DeletedAt); err != nil {
			return file, err
		}
	}

	if err = rows.Err(); err != nil {
		return file, err
	}

	return file, nil
}

// func (fileRepository *FileRepository) Update(file *entity.File) error {
// 	stmt, err := fileRepository.Db.Prepare("UPDATE files SET name = $1, picture = $2, is_deleted = $3, created_at = $4, updated_at = $5, deleted_at = $6 WHERE id = $7")
// 	if err != nil {
// 		return err
// 	}

// 	_, err = stmt.Exec(file.Name, file.Picture, file.IsDeleted, file.CreatedAt, file.UpdatedAt, file.DeletedAt, file.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (fileRepository *FileRepository) Delete(file *entity.File) error {
// 	stmt, err := fileRepository.Db.Prepare("UPDATE files SET is_deleted = $1, deleted_at = $2 WHERE id = $3")
// 	if err != nil {
// 		return err
// 	}

// 	_, err = stmt.Exec(file.IsDeleted, file.DeletedAt, file.ID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (fileRepository *FileRepository) IsDeleted(id string) error {
// 	var file entity.File

// 	rows, err := fileRepository.Db.Query("SELECT * FROM files WHERE id = $1", id)
// 	if err != nil {
// 		return err
// 	}

// 	for rows.Next() {
// 		if err := rows.Scan(&file.ID, &file.Name, &file.Picture, &file.IsDeleted, &file.CreatedAt, &file.UpdatedAt, &file.DeletedAt); err != nil {
// 			return err
// 		}
// 	}

// 	if err = rows.Err(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (fileRepository *FileRepository) FindAll() ([]entity.File, error) {
// 	rows, err := fileRepository.Db.Query("SELECT file_id, name, picture, is_deleted, created_at, updated_at, deleted_at FROM files")
// 	if err != nil {
// 		return nil, err
// 	}

// 	var files []entity.File

// 	for rows.Next() {
// 		var file entity.File

// 		if err := rows.Scan(&file.ID, &file.Name, &file.Picture, &file.IsDeleted, &file.CreatedAt, &file.UpdatedAt, &file.DeletedAt); err != nil {
// 			return files, err
// 		}

// 		files = append(files, file)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return files, err
// 	}

// 	return files, nil
// }
