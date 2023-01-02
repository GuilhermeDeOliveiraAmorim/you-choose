package main

import (
	"fmt"

	// actor "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/entity"
	chooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/entity"
	// director "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/director/entity"
	// genre "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/genre/entity"
	// movieList "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie-list/entity"
	// movie "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie/entity"
	// writer "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/writer/entity"
)

func main() {
	chooser, _ := chooser.NewChooser("Guilherme", "Amorim", "guiamorim", "guilherme.jpg", "asdqw2e23")

	// director, _ := director.NewDirector("Jose", "jose.jpg")

	fmt.Println(chooser)

	// actor, _ := actor.NewActor("Pedro", "pedro.jpg")

	// writer, _ := writer.NewWriter("Bob", "bob.jpg")

	// genre, _ := genre.NewGenre("acao", "acao.jpg")

	// movie, _ := movie.NewMovie("Filme Novo", "Like the previous output, your current date and time will be different from the example, but the format should be similar.", 4.8, "filme_novo.jpeg")

	// movie.AddActor(actor)
	// movie.AddDirector(director)
	// movie.AddWriter(writer)
	// movie.AddGenre(genre)

	// list, _ := movieList.NewMovieList("Nova Lista", "So you can print the current date and time in a format that’s")

	// list.AddMovie(movie)
	// list.AddChooser(chooser)

	// fmt.Println(list)
}
