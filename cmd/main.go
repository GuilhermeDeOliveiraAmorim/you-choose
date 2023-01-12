package main

import (
	"fmt"
	"time"

	// actor "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/actor/entity"
	// director "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/director/entity"
	// genre "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/genre/entity"
	// movieList "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie-list/entity"
	chooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/entity"
	"golang.org/x/crypto/bcrypt"
)

// writer "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/writer/entity"
// create_chooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/application/usecases/chooser/create_chooser"
// create_director "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/application/usecases/director/create_director"
// create_movie_list "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/application/usecases/movie-list/create_movie_list"

func main() {
	chooser, err := chooser.NewChooser("Guilherme", "Amorim", "guia", "guilherme", "AFT12õt$%#")

	// director, _ := director.NewDirector("Jose", "jose.jpg")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(chooser.ID)
		fmt.Println(chooser.FirstName)
		fmt.Println(chooser.LastName)
		fmt.Println(chooser.UserName)
		fmt.Println(chooser.Picture)
		fmt.Println(chooser.Password)
		fmt.Println(chooser.CreatedAt)
		fmt.Println(chooser.UpdatedAt)
		fmt.Println(chooser.IsDeleted)
	}

	currentTime := time.Now()

	fmt.Println(chooser.Picture + "-" + currentTime.Format("2006-01-02T15:04:05.000000Z") + ".jpg")

	// isTheSame := bcrypt.CompareHashAndPassword([]byte(chooser.Password), []byte("AFT12õts$%#"))
	// if isTheSame == nil {
	// 	fmt.Println("Ok")
	// } else {
	// 	fmt.Println(isTheSame)
	// }

	// inputDirector := &create_director.InputCreateDirectorDto{
	// 	Name:    "Guilherme",
	// 	Picture: "guilherme.jpg",
	// }

	// director := create_director.CreateDirectorUseCase(inputDirector)

	// fmt.Println(director)

	// inputMovieList := &create_movie_list.InputCreateMovieListDto{
	// 	Title:       "Nova Lista",
	// 	Description: "Uma ótima descrição",
	// 	Picture:     "nova_lista.jpg",
	// }

	// movieList := create_movie_list.CreateMovieList(inputMovieList)

	// fmt.Println(movieList)

	// inputChooser := &create_chooser.InputCreateChooserDto{
	// 	FirstName: "",
	// 	LastName:  "Amorim",
	// 	UserName:  "guiamorim123",
	// 	Picture:   "guilherme.jpg",
	// 	Password:  "AFT12õt$%#",
	// }

	// chooser := create_chooser.CreateChooserUseCase(inputChooser)

	// fmt.Println(chooser)

	// actor, _ := actor.NewActor("Pedro", "pedro.jpg")

	// writer, _ := writer.NewWriter("Bob", "bob.jpg")

	// genre, _ := genre.NewGenre("acao", "acao.jpg")

	// newMovie, _ := movie.NewMovie("Filme Novo", "Like the previous output, your current date and time will be different from the example, but the format should be similar.", 4.8, "filme_novo.jpeg")

	// fmt.Println(newMovie.GetVotes())

	// newMovie.AddVote()
	// newMovie.AddVote()
	// newMovie.AddVote()
	// newMovie.AddVote()
	// newMovie.AddVote()
	// newMovie.AddVote()
	// newMovie.AddVote()

	// movie.AddActor(actor)
	// movie.AddDirector(director)
	// movie.AddWriter(writer)
	// movie.AddGenre(genre)

	// list, _ := movieList.NewMovieList("Nova Lista", "So you can print the current date and time in a format that’s")

	// list.AddMovie(movie)
	// list.AddChooser(chooser)

	// fmt.Println(newMovie.GetVotes())

	// variavel := 10

	// fmt.Println(&variavel)

	// abc(&variavel)

	// fmt.Println(variavel)
	// fmt.Println(newMovie.GetVotes())
}

// Save password in hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// check password in hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
