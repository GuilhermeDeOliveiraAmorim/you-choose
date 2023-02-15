package main

import (
	"database/sql"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infra/web/webserver"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=root password=root dbname=root sslmode=disable")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	webserver := webserver.NewWebServer(":8080")
	newWebChooserHandler := NewWebChooserHandlerGen(db)
	newWebMovieListHandler := NewWebMovieListHandlerGen(db)
	newWebMovieHandler := NewWebMovieHandlerGen(db)
	newWebActorHandler := NewWebActorHandlerGen(db)
	newWebWriterHandler := NewWebWriterHandlerGen(db)
	newWebDirectorHandler := NewWebDirectorHandlerGen(db)
	newWebGenreHandler := NewWebGenreHandlerGen(db)

	webserver.AddHandler("/choosers/create/chooser", newWebChooserHandler.Create)
	webserver.AddHandler("/choosers/delete/chooser", newWebChooserHandler.Delete)
	webserver.AddHandler("/choosers/update/chooser", newWebChooserHandler.Update)
	webserver.AddHandler("/choosers/isdeleted/chooser", newWebChooserHandler.IsDeleted)
	webserver.AddHandler("/choosers/find/chooser", newWebChooserHandler.Find)
	webserver.AddHandler("/choosers/find/all/choosers", newWebChooserHandler.FindAll)
	webserver.AddHandler("/choosers/find/all/movielists", newWebChooserHandler.FindAllChooserMovieLists)
	webserver.AddHandler("/choosers/chooser/create/movielist", newWebChooserHandler.ChooserCreateMovieList)
	webserver.AddHandler("/choosers/chooser/add/movielist/movie", newWebChooserHandler.ChooserAddMovieToMovieList)

	webserver.AddHandler("/movielists/add", newWebMovieListHandler.Create)
	webserver.AddHandler("/movielists/all", newWebMovieListHandler.FindAll)
	webserver.AddHandler("/movielists/find", newWebMovieListHandler.Find)
	webserver.AddHandler("/movielists/add/chooser", newWebMovieListHandler.AddChooserToMovieList)

	webserver.AddHandler("/movies/create/movie", newWebMovieHandler.Create)
	webserver.AddHandler("/movies/find/movie", newWebMovieHandler.Find)
	webserver.AddHandler("/movies/find/all/movies", newWebMovieHandler.FindAll)

	webserver.AddHandler("/movies/add/movie/writers", newWebMovieHandler.AddWritersToMovie)
	webserver.AddHandler("/movies/find/movie/writers", newWebMovieHandler.FindMovieWriters)

	webserver.AddHandler("/movies/add/movie/actors", newWebMovieHandler.AddActorsToMovie)
	webserver.AddHandler("/movies/find/movie/actors", newWebMovieHandler.FindMovieActors)

	webserver.AddHandler("/movies/add/movie/directors", newWebMovieHandler.AddDirectorsToMovie)
	webserver.AddHandler("/movies/find/movie/directors", newWebMovieHandler.FindMovieDirectors)

	webserver.AddHandler("/movies/add/movie/genres", newWebMovieHandler.AddGenresToMovie)
	webserver.AddHandler("/movies/find/movie/genres", newWebMovieHandler.FindMovieGenres)

	webserver.AddHandler("/actors/create/actor", newWebActorHandler.Create)
	webserver.AddHandler("/actors/find/actor", newWebActorHandler.Find)
	webserver.AddHandler("/actors/find/all/actors", newWebActorHandler.FindAll)

	webserver.AddHandler("/writers/create/writer", newWebWriterHandler.Create)
	webserver.AddHandler("/writers/find/writer", newWebWriterHandler.Find)
	webserver.AddHandler("/writers/find/all/writers", newWebWriterHandler.FindAll)

	webserver.AddHandler("/directors/create/director", newWebDirectorHandler.Create)
	webserver.AddHandler("/directors/find/director", newWebDirectorHandler.Find)
	webserver.AddHandler("/directors/find/all/directors", newWebDirectorHandler.FindAll)

	webserver.AddHandler("/genres/create/genre", newWebGenreHandler.Create)
	webserver.AddHandler("/genres/find/genre", newWebGenreHandler.Find)
	webserver.AddHandler("/genres/find/all/genres", newWebGenreHandler.FindAll)

	webserver.Start()
}
