package database

import (
	"database/sql"

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
		return err
	}

	_, err = stmt.Exec(movieList.ID, movieList.Title, movieList.Description, movieList.Picture, movieList.IsDeleted, movieList.CreatedAt, movieList.UpdatedAt, movieList.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (movieListRepository *MovieListRepository) FindAll() ([]entity.MovieList, error) {
	rows, err := movieListRepository.Db.Query("SELECT id, title, description, picture, is_deleted, created_at, updated_at, deleted_at FROM movie_lists")
	if err != nil {
		return nil, err
	}

	var movieLists []entity.MovieList

	for rows.Next() {
		var movieList entity.MovieList

		if err := rows.Scan(&movieList.ID, &movieList.Title, &movieList.Description, &movieList.Picture, &movieList.IsDeleted, &movieList.CreatedAt, &movieList.UpdatedAt, &movieList.DeletedAt); err != nil {
			return movieLists, err
		}

		movieLists = append(movieLists, movieList)
	}

	if err = rows.Err(); err != nil {
		return movieLists, err
	}

	return movieLists, nil
}

func (movieListRepository *MovieListRepository) Find(id string) (entity.MovieList, error) {
	var movieList entity.MovieList

	rows, err := movieListRepository.Db.Query("SELECT * FROM movie_lists WHERE id = $1", id)
	if err != nil {
		return movieList, err
	}

	for rows.Next() {
		if err := rows.Scan(&movieList.ID, &movieList.Title, &movieList.Description, &movieList.Picture, &movieList.IsDeleted, &movieList.CreatedAt, &movieList.UpdatedAt, &movieList.DeletedAt); err != nil {
			return movieList, err
		}
	}

	if err = rows.Err(); err != nil {
		return movieList, err
	}

	return movieList, nil
}

func (movieListRepository *MovieListRepository) AddChooserToMovieList(movieList *entity.MovieList, chooser *entity.Chooser, created_at string, updated_at string, deleted_at string) error {
	stmt, err := movieListRepository.Db.Prepare("INSERT INTO choosers_movie_lists (chooser_id, movie_list_id, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(chooser.ID, movieList.ID, created_at, updated_at, deleted_at)
	if err != nil {
		return err
	}

	return nil
}