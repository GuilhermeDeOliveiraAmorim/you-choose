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

func (movieListRepository *MovieListRepository) Update(movieList *entity.MovieList) error {
	stmt, err := movieListRepository.Db.Prepare("UPDATE movie_lists SET title = $1, description = $2, picture = $3, is_deleted = $4, created_at = $5, updated_at = $6, deleted_at = $7 WHERE id = $8")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(movieList.Title, movieList.Description, movieList.Picture, movieList.IsDeleted, movieList.CreatedAt, movieList.UpdatedAt, movieList.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (movieListRepository *MovieListRepository) Delete(movieList *entity.MovieList) error {
	stmt, err := movieListRepository.Db.Prepare("UPDATE movie_lists SET is_deleted = $1, deleted_at = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(movieList.IsDeleted, movieList.DeletedAt, movieList.ID)
	if err != nil {
		return err
	}

	return nil
}

func (movieListRepository *MovieListRepository) IsDeleted(id string) error {
	var movieList entity.MovieList

	rows, err := movieListRepository.Db.Query("SELECT * FROM movieLists WHERE id = $1", id)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.Scan(
			&movieList.ID, &movieList.Title, &movieList.Description, &movieList.Picture, &movieList.IsDeleted, &movieList.CreatedAt, &movieList.UpdatedAt, &movieList.DeletedAt); err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
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

func (movieListRepository *MovieListRepository) FindMovieListMovies(id string) ([]string, error) {
	var moviesIds []string

	rows, err := movieListRepository.Db.Query("SELECT movie_id FROM movies_movie_lists WHERE movie_list_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var movieId string

		if err := rows.Scan(&movieId); err != nil {
			return moviesIds, err
		}

		moviesIds = append(moviesIds, movieId)
	}

	if err = rows.Err(); err != nil {
		return moviesIds, err
	}

	return moviesIds, nil
}
