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

	webserver.AddHandler("/choosers/create/chooser", newWebChooserHandler.Create)
	webserver.AddHandler("/choosers/delete/chooser", newWebChooserHandler.Delete)
	webserver.AddHandler("/choosers/update/chooser", newWebChooserHandler.Update)
	webserver.AddHandler("/choosers/isdeleted/chooser", newWebChooserHandler.IsDeleted)
	webserver.AddHandler("/choosers/find/chooser", newWebChooserHandler.Find)
	webserver.AddHandler("/choosers/find/all/choosers", newWebChooserHandler.FindAll)
	webserver.AddHandler("/choosers/find/all/movielists", newWebChooserHandler.FindAllChooserMovieLists)
	webserver.AddHandler("/choosers/chooser/create/movielist", newWebChooserHandler.ChooserCreateMovieList)
	webserver.AddHandler("/choosers/chooser/add/movie/movielist", newWebChooserHandler.ChooserAddMovieToMovieList)

	webserver.AddHandler("/movielists/add", newWebMovieListHandler.Create)
	webserver.AddHandler("/movielists/all", newWebMovieListHandler.FindAll)
	webserver.AddHandler("/movielists/find", newWebMovieListHandler.Find)
	webserver.AddHandler("/movielists/add/chooser", newWebMovieListHandler.AddChooserToMovieList)

	webserver.AddHandler("/movies/create/movie", newWebMovieHandler.Create)
	webserver.AddHandler("/movies/find/all/movies", newWebMovieHandler.FindAll)
	webserver.AddHandler("/movies/find/movie", newWebMovieHandler.Find)

	webserver.Start()
}
