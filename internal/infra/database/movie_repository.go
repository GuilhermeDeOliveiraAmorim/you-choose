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

func (movieRepository *MovieRepository) Create(movie *entity.Movie) error {
	stmt, err := movieRepository.Db.Prepare("INSERT INTO movies (movie_id, title, synopsis, imdb_rating, votes, you_choose_rating, poster, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(movie.ID, movie.Title, movie.Synopsis, movie.ImdbRating, movie.Votes, movie.YouChooseRating, movie.Poster, movie.IsDeleted, movie.CreatedAt, movie.UpdatedAt, movie.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (movieRepository *MovieRepository) FindAll() ([]entity.Movie, error) {
	rows, err := movieRepository.Db.Query("SELECT movie_id, title, synopsis, imdb_rating, votes, you_choose_rating, poster, is_deleted, created_at, updated_at, deleted_at FROM movies")
	if err != nil {
		return nil, err
	}

	var movies []entity.Movie

	for rows.Next() {
		var movie entity.Movie

		if err := rows.Scan(&movie.ID, &movie.Title, &movie.Synopsis, &movie.ImdbRating, &movie.Votes, &movie.YouChooseRating, &movie.Poster, &movie.IsDeleted, &movie.CreatedAt, &movie.UpdatedAt, &movie.DeletedAt); err != nil {
			return movies, err
		}

		movies = append(movies, movie)
	}

	if err = rows.Err(); err != nil {
		return movies, err
	}

	return movies, nil
}

func (movieRepository *MovieRepository) Find(id string) (entity.Movie, error) {
	var movie entity.Movie

	rows, err := movieRepository.Db.Query("SELECT * FROM movies WHERE movie_id = $1", id)
	if err != nil {
		return movie, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Synopsis,
			&movie.ImdbRating,
			&movie.Votes,
			&movie.YouChooseRating,
			&movie.Poster,
			&movie.IsDeleted,
			&movie.CreatedAt,
			&movie.UpdatedAt,
			&movie.DeletedAt); err != nil {
			return movie, err
		}
	}

	if err = rows.Err(); err != nil {
		return movie, err
	}

	return movie, nil
}

func (movieRepository *MovieRepository) AddActorsToMovie(movie entity.Movie, actors []*entity.Actor) error {
	for _, actor := range movie.Actors {
		stmt, err := movieRepository.Db.Prepare("INSERT INTO actors_movies (movie_id, actor_id, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(&movie.ID, &actor.ID, false, &movie.UpdatedAt, &movie.UpdatedAt, &movie.UpdatedAt)
		if err != nil {
			return err
		}

	}

	return nil
}

func (movieRepository *MovieRepository) FindMovieActors(id string) ([]entity.Actor, error) {
	var actors []entity.Actor

	return actors, nil
}
