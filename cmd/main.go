package main

import (
	"fmt"

	director "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/director/entity"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/entity"
)

func main() {
	chooser := entity.NewChooser("Guilherme", "Amorim", "guiamorim", "guilherme.jpg")

	director := director.NewDirector("Jose", "jose.jpg")

	actor := entity.NewActor("Pedro", "pedro.jpg")

	writer := entity.NewWriter("Bob", "bob.jpg")

	genre := entity.NewGenre("acao", "acao.jpg")

	movie := entity.NewMovie("Filme Novo", "Like the previous output, your current date and time will be different from the example, but the format should be similar.", 4.8, "filme_novo.jpeg")

	movie.AddActor(actor)
	movie.AddDirector(director)
	movie.AddWriter(writer)
	movie.AddGenre(genre)

	list := entity.NewMovieList("Nova Lista", "So you can print the current date and time in a format that’s")

	list.AddMovie(movie)
	list.AddChooser(chooser)

	fmt.Println(list)
}
