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
	newWebTaskHandler := NewWebChooserHandlerGen(db)

	webserver.AddHandler("/choosers/add", newWebTaskHandler.Create)
	webserver.AddHandler("/choosers", newWebTaskHandler.FindAll)
	webserver.AddHandler("/chooser", newWebTaskHandler.Find)

	webserver.Start()
}
