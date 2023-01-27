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

	webserver.AddHandler("/choosers/add", newWebChooserHandler.Create)
	webserver.AddHandler("/choosers/find", newWebChooserHandler.Find)
	webserver.AddHandler("/choosers/delete", newWebChooserHandler.Delete)
	webserver.AddHandler("/choosers/update", newWebChooserHandler.Update)
	webserver.AddHandler("/choosers/all", newWebChooserHandler.FindAll)
	webserver.AddHandler("/choosers/isdeleted", newWebChooserHandler.IsDeleted)

	webserver.AddHandler("/movielists/add", newWebMovieListHandler.Create)
	webserver.AddHandler("/movielists", newWebMovieListHandler.FindAll)
	webserver.AddHandler("/movielist", newWebMovieListHandler.Find)

	webserver.Start()
}
