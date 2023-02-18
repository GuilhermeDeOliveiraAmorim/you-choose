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
	newWebTagHandler := NewWebTagHandlerGen(db)

	webserver.AddHandler("/choosers/create/chooser", newWebChooserHandler.Create)
	webserver.AddHandler("/choosers/find/chooser", newWebChooserHandler.Find)
	webserver.AddHandler("/choosers/update/chooser", newWebChooserHandler.Update)
	webserver.AddHandler("/choosers/delete/chooser", newWebChooserHandler.Delete)
	webserver.AddHandler("/choosers/isdeleted/chooser", newWebChooserHandler.IsDeleted)
	webserver.AddHandler("/choosers/find/all/choosers", newWebChooserHandler.FindAll)
	webserver.AddHandler("/choosers/add/movies/movielist", newWebChooserHandler.AddMoviesToMovieList)
	webserver.AddHandler("/choosers/add/choosers/movielist", newWebChooserHandler.AddChoosersToMovieList)
	webserver.AddHandler("/choosers/add/tags/movielist", newWebChooserHandler.AddTagsToMovieList)

	webserver.AddHandler("/movielists/create/movielist", newWebMovieListHandler.Create)
	webserver.AddHandler("/movielists/find/movielist", newWebMovieListHandler.Find)
	webserver.AddHandler("/movielists/update/movielist", newWebMovieListHandler.Update)
	webserver.AddHandler("/movielists/delete/movielist", newWebMovieListHandler.Delete)
	webserver.AddHandler("/movielists/isdeleted/movielist", newWebMovieListHandler.IsDeleted)
	webserver.AddHandler("/movielists/find/all/movielists", newWebMovieListHandler.FindAll)
	webserver.AddHandler("/movielists/find/movielist/movies", newWebMovieListHandler.FindMovieListMovies)
	webserver.AddHandler("/movielists/find/movielist/choosers", newWebMovieListHandler.FindMovieListChoosers)
	webserver.AddHandler("/movielists/find/movielist/tags", newWebMovieListHandler.FindMovieListTags)

	webserver.AddHandler("/movies/create/movie", newWebMovieHandler.Create)
	webserver.AddHandler("/movies/find/movie", newWebMovieHandler.Find)
	webserver.AddHandler("/movies/update/movie", newWebMovieHandler.Update)
	webserver.AddHandler("/movies/delete/movie", newWebMovieHandler.Delete)
	webserver.AddHandler("/movies/isdeleted/movie", newWebMovieHandler.IsDeleted)
	webserver.AddHandler("/movies/find/all/movies", newWebMovieHandler.FindAll)
	webserver.AddHandler("/movies/add/vote/movie", newWebMovieHandler.AddVoteToMovie)

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
	webserver.AddHandler("/actors/update/actor", newWebActorHandler.Update)
	webserver.AddHandler("/actors/delete/actor", newWebActorHandler.Delete)
	webserver.AddHandler("/actors/isdeleted/actor", newWebActorHandler.IsDeleted)
	webserver.AddHandler("/actors/find/all/actors", newWebActorHandler.FindAll)

	webserver.AddHandler("/writers/create/writer", newWebWriterHandler.Create)
	webserver.AddHandler("/writers/find/writer", newWebWriterHandler.Find)
	webserver.AddHandler("/writers/update/writer", newWebWriterHandler.Update)
	webserver.AddHandler("/writers/delete/writer", newWebWriterHandler.Delete)
	webserver.AddHandler("/writers/isdeleted/writer", newWebWriterHandler.IsDeleted)
	webserver.AddHandler("/writers/find/all/writers", newWebWriterHandler.FindAll)

	webserver.AddHandler("/directors/create/director", newWebDirectorHandler.Create)
	webserver.AddHandler("/directors/find/director", newWebDirectorHandler.Find)
	webserver.AddHandler("/directors/update/director", newWebDirectorHandler.Update)
	webserver.AddHandler("/directors/delete/director", newWebDirectorHandler.Delete)
	webserver.AddHandler("/directors/isdeleted/director", newWebDirectorHandler.IsDeleted)
	webserver.AddHandler("/directors/find/all/directors", newWebDirectorHandler.FindAll)

	webserver.AddHandler("/genres/create/genre", newWebGenreHandler.Create)
	webserver.AddHandler("/genres/find/genre", newWebGenreHandler.Find)
	webserver.AddHandler("/genres/update/genre", newWebGenreHandler.Update)
	webserver.AddHandler("/genres/delete/genre", newWebGenreHandler.Delete)
	webserver.AddHandler("/genres/isdeleted/genre", newWebGenreHandler.IsDeleted)
	webserver.AddHandler("/genres/find/all/genres", newWebGenreHandler.FindAll)

	webserver.AddHandler("/tags/create/tag", newWebTagHandler.Create)
	webserver.AddHandler("/tags/find/tag", newWebTagHandler.Find)
	webserver.AddHandler("/tags/update/tag", newWebTagHandler.Update)
	webserver.AddHandler("/tags/delete/tag", newWebTagHandler.Delete)
	webserver.AddHandler("/tags/isdeleted/tag", newWebTagHandler.IsDeleted)
	webserver.AddHandler("/tags/find/all/tags", newWebTagHandler.FindAll)

	webserver.Start()
}
