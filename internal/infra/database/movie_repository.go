package database

import (
	"database/sql"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type MovieRepository struct {
	Db *sql.DB
}

func NewMovieRepository(db *sql.DB) *MovieRepository {
	return &MovieRepository{
		Db: db,
	}
}

func (c *MovieRepository) Create(movie *entity.Movie) error {
	stmt, err := c.Db.Prepare("INSERT INTO movies (movie_id, title, synopsis, imdb_rating, votes, you_choose_rating, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(movie.ID, movie.Title, movie.Synopsis, movie.ImdbRating, movie.Votes, movie.YouChooseRating, movie.CreatedAt, movie.UpdatedAt, movie.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (c *MovieRepository) FindAll() ([]entity.Movie, error) {
	rows, err := c.Db.Query("SELECT movie_id, title, synopsis, imdb_rating, votes, you_choose_rating, is_deleted, created_at, updated_at, deleted_at FROM movies")
	if err != nil {
		return nil, err
	}

	var movies []entity.Movie

	for rows.Next() {
		var movie entity.Movie

		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Synopsis, &movie.ImdbRating, &movie.Votes, &movie.YouChooseRating, &movie.IsDeleted, &movie.CreatedAt, &movie.UpdatedAt, &movie.DeletedAt); err != nil {
			return movies, err
		}

		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		return movies, err
	}

	return movies, nil
}
