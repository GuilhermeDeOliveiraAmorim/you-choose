package main

import (
	"fmt"
	// actor "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/entity"
	// director "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/director/entity"
	// genre "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/genre/entity"
	// movieList "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie-list/entity"
	// movie "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie/entity"
	// writer "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/writer/entity"

	create_chooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/application/usecases/chooser/create_chooser"
	create_director "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/application/usecases/director/create_director"
	add_chooser_to_list "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/application/usecases/movie-list/add_chooser_to_list"
	create_movie_list "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/application/usecases/movie-list/create_movie_list"
	c "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/entity"
	ml "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie-list/entity"
)

func main() {
	// chooser, _ := chooser.NewChooser("Guilherme", "Amorim", "guiamorim", "guilherme.jpg", "asdqw2e23")

	// director, _ := director.NewDirector("Jose", "jose.jpg")

	// fmt.Println(chooser)

	inputDirector := &create_director.InputCreateDirectorDto{
		Name:    "Guilherme",
		Picture: "guilherme.jpg",
	}

	director := create_director.CreateDirectorUseCase(inputDirector)

	fmt.Println(director)

	inputMovieList := &create_movie_list.InputCreateMovieListDto{
		Title:       "Nova Lista",
		Description: "Uma ótima descrição",
		Picture:     "nova_lista.jpg",
	}

	movieList := create_movie_list.CreateMovieList(inputMovieList)

	fmt.Println(movieList)

	inputChooser := &create_chooser.InputCreateChooserDto{
		FirstName: "Guilherme",
		LastName:  "Amorim",
		UserName:  "guiamorim",
		Picture:   "guilherme.jpg",
		Password:  "asd123",
	}

	chooser := create_chooser.CreateChooserUseCase(inputChooser)

	inputAddChooserToMovieList := &add_chooser_to_list.InputAddChooserToMovieListDto{
		MovieList: &ml.MovieList{ID: movieList.ID, Title: movieList.Title, Description: movieList.Description},
		Chooser:   &c.Chooser{ID: chooser.ID, UserName: chooser.UserName},
	}

	add_chooser_to_list.AddChooserToList(inputAddChooserToMovieList)

	fmt.Println(movieList)

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
