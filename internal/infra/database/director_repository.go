package database

import (
	"database/sql"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type DirectorRepository struct {
	Db *sql.DB
}

func NewDirectorRepository(db *sql.DB) *DirectorRepository {
	return &DirectorRepository{
		Db: db,
	}
}

func (directorRepository *DirectorRepository) Create(director *entity.Director) error {
	stmt, err := directorRepository.Db.Prepare("INSERT INTO directors (director_id, name, picture, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(director.ID, director.Name, director.Picture, director.IsDeleted, director.CreatedAt, director.UpdatedAt, director.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (directorRepository *DirectorRepository) Find(id string) (entity.Director, error) {
	var director entity.Director

	rows, err := directorRepository.Db.Query("SELECT * FROM directors WHERE director_id = $1", id)
	if err != nil {
		return director, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&director.ID,
			&director.Name,
			&director.Picture,
			&director.IsDeleted,
			&director.CreatedAt,
			&director.UpdatedAt,
			&director.DeletedAt); err != nil {
			return director, err
		}
	}

	if err = rows.Err(); err != nil {
		return director, err
	}

	return director, nil
}

func (directorRepository *DirectorRepository) Update(director *entity.Director) error {
	stmt, err := directorRepository.Db.Prepare("UPDATE directors SET name = $1, picture = $2, is_deleted = $3, created_at = $4, updated_at = $5, deleted_at = $6 WHERE id = $7")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(director.Name, director.Picture, director.IsDeleted, director.CreatedAt, director.UpdatedAt, director.DeletedAt, director.ID)
	if err != nil {
		return err
	}

	return nil
}

func (directorRepository *DirectorRepository) Delete(director *entity.Director) error {
	stmt, err := directorRepository.Db.Prepare("UPDATE directors SET is_deleted = $1, deleted_at = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(director.IsDeleted, director.DeletedAt, director.ID)
	if err != nil {
		return err
	}

	return nil
}

func (directorRepository *DirectorRepository) IsDeleted(id string) error {
	var director entity.Director

	rows, err := directorRepository.Db.Query("SELECT * FROM directors WHERE id = $1", id)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.Scan(&director.ID, &director.Name, &director.Picture, &director.IsDeleted, &director.CreatedAt, &director.UpdatedAt, &director.DeletedAt); err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (directorRepository *DirectorRepository) FindAll() ([]entity.Director, error) {
	rows, err := directorRepository.Db.Query("SELECT director_id, name, picture, is_deleted, created_at, updated_at, deleted_at FROM directors")
	if err != nil {
		return nil, err
	}

	var directors []entity.Director

	for rows.Next() {
		var director entity.Director

		if err := rows.Scan(&director.ID, &director.Name, &director.Picture, &director.IsDeleted, &director.CreatedAt, &director.UpdatedAt, &director.DeletedAt); err != nil {
			return directors, err
		}

		directors = append(directors, director)
	}

	if err = rows.Err(); err != nil {
		return directors, err
	}

	return directors, nil
}
