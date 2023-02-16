package database

import (
	"database/sql"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type ChooserRepository struct {
	Db *sql.DB
}

func NewChooserRepository(db *sql.DB) *ChooserRepository {
	return &ChooserRepository{
		Db: db,
	}
}

func (chooserRepository *ChooserRepository) Create(chooser *entity.Chooser) error {
	stmt, err := chooserRepository.Db.Prepare("INSERT INTO choosers (id, first_name, last_name, username, picture, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(chooser.ID, chooser.FirstName, chooser.LastName, chooser.UserName, chooser.Picture, chooser.IsDeleted, chooser.CreatedAt, chooser.UpdatedAt, chooser.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (chooserRepository *ChooserRepository) Find(id string) (entity.Chooser, error) {
	var chooser entity.Chooser

	rows, err := chooserRepository.Db.Query("SELECT * FROM choosers WHERE id = $1", id)
	if err != nil {
		return chooser, err
	}

	for rows.Next() {
		if err := rows.Scan(&chooser.ID, &chooser.FirstName, &chooser.LastName, &chooser.UserName, &chooser.Picture, &chooser.IsDeleted, &chooser.CreatedAt, &chooser.UpdatedAt, &chooser.DeletedAt); err != nil {
			return chooser, err
		}
	}

	if err = rows.Err(); err != nil {
		return chooser, err
	}

	return chooser, nil
}

func (chooserRepository *ChooserRepository) Update(chooser *entity.Chooser) error {
	stmt, err := chooserRepository.Db.Prepare("UPDATE choosers SET first_name = $1, last_name = $2, username = $3, picture = $4, updated_at = $5 WHERE id = $6")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(chooser.FirstName, chooser.LastName, chooser.UserName, chooser.Picture, chooser.UpdatedAt, chooser.ID)
	if err != nil {
		return err
	}

	return nil
}

func (chooserRepository *ChooserRepository) Delete(chooser *entity.Chooser) error {
	stmt, err := chooserRepository.Db.Prepare("UPDATE choosers SET is_deleted = $1, deleted_at = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(chooser.IsDeleted, chooser.DeletedAt, chooser.ID)
	if err != nil {
		return err
	}

	return nil
}

func (chooserRepository *ChooserRepository) IsDeleted(id string) error {
	var chooser entity.Chooser

	rows, err := chooserRepository.Db.Query("SELECT * FROM choosers WHERE id = $1", id)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.Scan(&chooser.ID, &chooser.FirstName, &chooser.LastName, &chooser.UserName, &chooser.Picture, &chooser.IsDeleted, &chooser.CreatedAt, &chooser.UpdatedAt, &chooser.DeletedAt); err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (chooserRepository *ChooserRepository) FindAll() ([]entity.Chooser, error) {
	rows, err := chooserRepository.Db.Query("SELECT id, first_name, last_name, username, picture, is_deleted, created_at, updated_at, deleted_at FROM choosers")
	if err != nil {
		return nil, err
	}

	var choosers []entity.Chooser

	for rows.Next() {
		var chooser entity.Chooser

		if err := rows.Scan(&chooser.ID, &chooser.FirstName, &chooser.LastName, &chooser.UserName, &chooser.Picture, &chooser.IsDeleted, &chooser.CreatedAt, &chooser.UpdatedAt, &chooser.DeletedAt); err != nil {
			return choosers, err
		}

		choosers = append(choosers, chooser)
	}

	if err = rows.Err(); err != nil {
		return choosers, err
	}

	return choosers, nil
}

func (chooserRepository *ChooserRepository) AddMoviesToMovieList(movieList entity.MovieList, movies []entity.Movie) error {
	for _, movie := range movies {
		stmt, err := chooserRepository.Db.Prepare("INSERT INTO movies_movie_lists (movie_id, movie_list_id, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(&movie.ID, &movieList.ID, false, &movieList.UpdatedAt, &movieList.UpdatedAt, &movieList.UpdatedAt)
		if err != nil {
			return err
		}

	}

	return nil
}

func (chooserRepository *ChooserRepository) AddChoosersToMovieList(movieList entity.MovieList, choosers []entity.Chooser) error {
	for _, chooser := range choosers {
		stmt, err := chooserRepository.Db.Prepare("INSERT INTO choosers_movie_lists (chooser_id, movie_list_id, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(&chooser.ID, &movieList.ID, false, &movieList.UpdatedAt, &movieList.UpdatedAt, &movieList.UpdatedAt)
		if err != nil {
			return err
		}

	}

	return nil
}
