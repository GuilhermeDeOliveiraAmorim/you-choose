package main

import (
	"fmt"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/configs"
	repo "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/repository"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infra/web/webserver"
)

type MorimServer struct {
	Repo repo.ActorRepositoryInterface
}

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	webserver := webserver.NewWebServer(configs.WebServerPort)
	// webActorHandler := web.NewWebActorHandler(MorimServer)
	// webserver.AddHandler("/actor", webActorHandler.Create)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()
}
