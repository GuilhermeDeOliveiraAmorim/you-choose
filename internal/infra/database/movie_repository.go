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

func (movieRepository *MovieRepository) Update(movie *entity.Movie) error {
	stmt, err := movieRepository.Db.Prepare("UPDATE movies SET title = $1, synopsis = $2, imdb_rating = $3, votes = $4, you_choose_rating = $5, poster = $6, is_deleted = $7, created_at = $8, updated_at = $9, deleted_at = $10 WHERE id = $11")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(movie.Title, movie.Synopsis, movie.ImdbRating, movie.Votes, movie.YouChooseRating, movie.Poster, movie.IsDeleted, movie.CreatedAt, movie.UpdatedAt, movie.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (movieRepository *MovieRepository) Delete(movie *entity.Movie) error {
	stmt, err := movieRepository.Db.Prepare("UPDATE movies SET is_deleted = $1, deleted_at = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(movie.IsDeleted, movie.DeletedAt, movie.ID)
	if err != nil {
		return err
	}

	return nil
}

func (movieRepository *MovieRepository) IsDeleted(id string) error {
	var movie entity.Movie

	rows, err := movieRepository.Db.Query("SELECT * FROM movies WHERE id = $1", id)
	if err != nil {
		return err
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
			return err
		}
	}

	if err = rows.Err(); err != nil {
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

func (movieRepository *MovieRepository) AddActorsToMovie(movie entity.Movie, actors []entity.Actor) error {
	for _, actor := range actors {
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

func (movieRepository *MovieRepository) FindMovieActors(id string) ([]string, error) {
	var actorsIds []string

	rows, err := movieRepository.Db.Query("SELECT actor_id FROM actors_movies WHERE movie_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var actorId string

		if err := rows.Scan(&actorId); err != nil {
			return actorsIds, err
		}

		actorsIds = append(actorsIds, actorId)
	}

	if err = rows.Err(); err != nil {
		return actorsIds, err
	}

	return actorsIds, nil
}

func (movieRepository *MovieRepository) AddWritersToMovie(movie entity.Movie, writers []entity.Writer) error {
	for _, writer := range writers {
		stmt, err := movieRepository.Db.Prepare("INSERT INTO writers_movies (movie_id, writer_id, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(&movie.ID, &writer.ID, false, &movie.UpdatedAt, &movie.UpdatedAt, &movie.UpdatedAt)
		if err != nil {
			return err
		}

	}

	return nil
}

func (movieRepository *MovieRepository) FindMovieWriters(id string) ([]string, error) {
	var writersIds []string

	rows, err := movieRepository.Db.Query("SELECT writer_id FROM writers_movies WHERE movie_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var writerId string

		if err := rows.Scan(&writerId); err != nil {
			return writersIds, err
		}

		writersIds = append(writersIds, writerId)
	}

	if err = rows.Err(); err != nil {
		return writersIds, err
	}

	return writersIds, nil
}

func (movieRepository *MovieRepository) AddDirectorsToMovie(movie entity.Movie, directors []entity.Director) error {
	for _, director := range directors {
		stmt, err := movieRepository.Db.Prepare("INSERT INTO directors_movies (movie_id, director_id, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(&movie.ID, &director.ID, false, &movie.UpdatedAt, &movie.UpdatedAt, &movie.UpdatedAt)
		if err != nil {
			return err
		}

	}

	return nil
}

func (movieRepository *MovieRepository) FindMovieDirectors(id string) ([]string, error) {
	var directorsIds []string

	rows, err := movieRepository.Db.Query("SELECT director_id FROM directors_movies WHERE movie_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var directorId string

		if err := rows.Scan(&directorId); err != nil {
			return directorsIds, err
		}

		directorsIds = append(directorsIds, directorId)
	}

	if err = rows.Err(); err != nil {
		return directorsIds, err
	}

	return directorsIds, nil
}

func (movieRepository *MovieRepository) AddGenresToMovie(movie entity.Movie, genres []entity.Genre) error {
	for _, genre := range genres {
		stmt, err := movieRepository.Db.Prepare("INSERT INTO genres_movies (movie_id, genre_id, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(&movie.ID, &genre.ID, false, &movie.UpdatedAt, &movie.UpdatedAt, &movie.UpdatedAt)
		if err != nil {
			return err
		}

	}

	return nil
}

func (movieRepository *MovieRepository) FindMovieGenres(id string) ([]string, error) {
	var genresIds []string

	rows, err := movieRepository.Db.Query("SELECT genre_id FROM genres_movies WHERE movie_id = $1", id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var genreId string

		if err := rows.Scan(&genreId); err != nil {
			return genresIds, err
		}

		genresIds = append(genresIds, genreId)
	}

	if err = rows.Err(); err != nil {
		return genresIds, err
	}

	return genresIds, nil
}
