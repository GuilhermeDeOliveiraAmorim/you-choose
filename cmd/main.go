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

	webserver.AddHandler("/choosers/create/chooser", newWebChooserHandler.Create, "Post")
	webserver.AddHandler("/choosers/delete/chooser", newWebChooserHandler.Delete, "Delete")
	webserver.AddHandler("/choosers/update/chooser", newWebChooserHandler.Update, "Put")
	webserver.AddHandler("/choosers/isdeleted/chooser", newWebChooserHandler.IsDeleted, "Delete")
	webserver.AddHandler("/choosers/find/chooser", newWebChooserHandler.Find, "Get")
	webserver.AddHandler("/choosers/find/all/choosers", newWebChooserHandler.FindAll, "Get")
	webserver.AddHandler("/choosers/find/all/movielists", newWebChooserHandler.FindAllChooserMovieLists, "Get")
	webserver.AddHandler("/choosers/chooser/create/movielist", newWebChooserHandler.ChooserCreateMovieList, "Post")
	webserver.AddHandler("/choosers/chooser/add/movie/movielist", newWebChooserHandler.ChooserAddMovieToMovieList, "Post")

	webserver.AddHandler("/movielists/add", newWebMovieListHandler.Create, "Post")
	webserver.AddHandler("/movielists/all", newWebMovieListHandler.FindAll, "Get")
	webserver.AddHandler("/movielists/find", newWebMovieListHandler.Find, "Get")
	webserver.AddHandler("/movielists/add/chooser", newWebMovieListHandler.AddChooserToMovieList, "Post")

	webserver.AddHandler("/movies/create/movie", newWebMovieHandler.Create, "Post")
	webserver.AddHandler("/movies/find/all/movies", newWebMovieHandler.FindAll, "Get")
	webserver.AddHandler("/movies/find/movie", newWebMovieHandler.Find, "Get")

	webserver.Start()
}
