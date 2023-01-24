package database

import (
	"database/sql"
	"fmt"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type MovieListRepository struct {
	Db *sql.DB
}

func NewMovieListRepository(db *sql.DB) *MovieListRepository {
	return &MovieListRepository{
		Db: db,
	}
}

func (movieListRepository *MovieListRepository) Create(movieList *entity.MovieList) error {
	stmt, err := movieListRepository.Db.Prepare("INSERT INTO movie_lists (id, title, description, picture, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		fmt.Print(err)
		return err
	}
	_, err = stmt.Exec(movieList.ID, movieList.Title, movieList.Description, movieList.Picture, movieList.IsDeleted, movieList.CreatedAt, movieList.UpdatedAt, movieList.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

// func (c *ChooserRepository) FindAll() ([]entity.Chooser, error) {
// 	rows, err := c.Db.Query("SELECT id, first_name, last_name, username, password, picture, is_deleted, created_at, updated_at, deleted_at FROM choosers")
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil, err
// 	}

// 	var choosers []entity.Chooser

// 	for rows.Next() {
// 		var chooser entity.Chooser

// 		if err := rows.Scan(&chooser.ID, &chooser.FirstName, &chooser.LastName, &chooser.UserName, &chooser.Password, &chooser.Picture, &chooser.IsDeleted, &chooser.CreatedAt, &chooser.UpdatedAt, &chooser.DeletedAt); err != nil {
// 			return choosers, err
// 		}

// 		choosers = append(choosers, chooser)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return choosers, err
// 	}

// 	return choosers, nil
// }

// func (chooserRepository *ChooserRepository) Find(id string) (entity.Chooser, error) {
// 	var chooser entity.Chooser

// 	rows, err := chooserRepository.Db.Query("SELECT * FROM choosers WHERE id = $1", id)
// 	if err != nil {
// 		return chooser, err
// 	}

// 	if err := rows.Scan(&chooser); err != nil {
// 		return chooser, err
// 	}

// 	if err = rows.Err(); err != nil {
// 		return chooser, err
// 	}

// 	return chooser, nil
// }
