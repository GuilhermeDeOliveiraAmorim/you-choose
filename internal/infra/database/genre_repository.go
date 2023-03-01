package database

import (
	"database/sql"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type GenreRepository struct {
	Db *sql.DB
}

func NewGenreRepository(db *sql.DB) *GenreRepository {
	return &GenreRepository{
		Db: db,
	}
}

func (genreRepository *GenreRepository) Create(genre *entity.Genre) error {
	stmt, err := genreRepository.Db.Prepare("INSERT INTO genres (genre_id, name, picture, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(genre.ID, genre.Name, genre.Picture, genre.IsDeleted, genre.CreatedAt, genre.UpdatedAt, genre.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (genreRepository *GenreRepository) Find(id string) (entity.Genre, error) {
	var genre entity.Genre

	rows, err := genreRepository.Db.Query("SELECT * FROM genres WHERE genre_id = $1", id)
	if err != nil {
		return genre, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&genre.ID,
			&genre.Name,
			&genre.Picture,
			&genre.IsDeleted,
			&genre.CreatedAt,
			&genre.UpdatedAt,
			&genre.DeletedAt); err != nil {
			return genre, err
		}
	}

	if err = rows.Err(); err != nil {
		return genre, err
	}

	return genre, nil
}

func (genreRepository *GenreRepository) Update(genre *entity.Genre) error {
	stmt, err := genreRepository.Db.Prepare("UPDATE genres SET name = $1, picture = $2, is_deleted = $3, created_at = $4, updated_at = $5, deleted_at = $6 WHERE genre_id = $7")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(genre.Name, genre.Picture, genre.IsDeleted, genre.CreatedAt, genre.UpdatedAt, genre.DeletedAt, genre.ID)
	if err != nil {
		return err
	}

	return nil
}

func (genreRepository *GenreRepository) Delete(genre *entity.Genre) error {
	stmt, err := genreRepository.Db.Prepare("UPDATE genres SET is_deleted = $1, deleted_at = $2 WHERE genre_id = $3")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(genre.IsDeleted, genre.DeletedAt, genre.ID)
	if err != nil {
		return err
	}

	return nil
}

func (genreRepository *GenreRepository) IsDeleted(id string) error {
	var genre entity.Genre

	rows, err := genreRepository.Db.Query("SELECT * FROM genres WHERE genre_id = $1", id)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.Scan(&genre.ID, &genre.Name, &genre.Picture, &genre.IsDeleted, &genre.CreatedAt, &genre.UpdatedAt, &genre.DeletedAt); err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (genreRepository *GenreRepository) FindAll() ([]entity.Genre, error) {
	rows, err := genreRepository.Db.Query("SELECT genre_id, name, picture, is_deleted, created_at, updated_at, deleted_at FROM genres")
	if err != nil {
		return nil, err
	}

	var genres []entity.Genre

	for rows.Next() {
		var genre entity.Genre

		if err := rows.Scan(&genre.ID, &genre.Name, &genre.Picture, &genre.IsDeleted, &genre.CreatedAt, &genre.UpdatedAt, &genre.DeletedAt); err != nil {
			return genres, err
		}

		genres = append(genres, genre)
	}

	if err = rows.Err(); err != nil {
		return genres, err
	}

	return genres, nil
}

func (genreRepository *GenreRepository) FindGenreByName(name string) (entity.Genre, error) {
	var genre entity.Genre

	rows, err := genreRepository.Db.Query("SELECT * FROM genres WHERE name = $1", name)
	if err != nil {
		return genre, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&genre.ID,
			&genre.Name,
			&genre.Picture,
			&genre.IsDeleted,
			&genre.CreatedAt,
			&genre.UpdatedAt,
			&genre.DeletedAt); err != nil {
			return genre, err
		}
	}

	if err = rows.Err(); err != nil {
		return genre, err
	}

	return genre, nil
}
