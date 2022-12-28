package main

import (
	"fmt"
	"time"
)

type Chooser struct {
	Id        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Picture   string    `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
}

type Movie struct {
	Id         string     `json:"id"`
	Title      string     `json:"title"`
	Synopsis   string     `json:"synopsis"`
	ImdbRating float32    `json:"imdb_rating"`
	Poster     string     `json:"poster"`
	Directors  []Director `json:"directors"`
	Actors     []Actor    `json:"actors"`
	Writers    []Writer   `json:"writers"`
	Genres     []Genre    `json:"genres"`
	CreatedAt  time.Time  `json:"created_at"`
	EditedAt   time.Time  `json:"edited_at"`
}

type List struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Chooser     Chooser   `json:"chooser"`
	Movies      []Movie   `json:"movies"`
	CreatedAt   time.Time `json:"created_at"`
	EditedAt    time.Time `json:"edited_at"`
}

type Director struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Picture   string    `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
}

type Actor struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Picture   string    `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
}

type Writer struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Picture   string    `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
}

type Genre struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Picture   string    `json:"picture"`
	CreatedAt time.Time `json:"created_at"`
	EditedAt  time.Time `json:"edited_at"`
}

func main() {
	chooser := Chooser{
		Id:        "asd",
		FirstName: "Guilherme",
		LastName:  "Amorim",
		Picture:   "guilherme.jpg",
		CreatedAt: time.Now(),
	}

	director := Director{
		Id:        "qwe",
		Name:      "José",
		Picture:   "jose.jpg",
		CreatedAt: time.Now(),
	}

	actor := Actor{
		Id:        "tyu",
		Name:      "Pedro",
		Picture:   "pedro.jpg",
		CreatedAt: time.Now(),
	}

	writer := Writer{
		Id:        "uiop",
		Name:      "Bob",
		Picture:   "bob.jpg",
		CreatedAt: time.Now(),
	}

	genre := Genre{
		Id:        "hjjr",
		Name:      "ação",
		Picture:   "acao.jpg",
		CreatedAt: time.Now(),
	}

	movie := Movie{
		Id:         "poih",
		Title:      "Filme Novo",
		Synopsis:   "Like the previous output, your current date and time will be different from the example, but the format should be similar.",
		ImdbRating: 4.8,
		Poster:     "filme_novo.jpeg",
		Directors:  []Director{director},
		Actors:     []Actor{actor},
		Writers:    []Writer{writer},
		Genres:     []Genre{genre},
		CreatedAt:  time.Now(),
		EditedAt:   time.Now(),
	}

	list := List{
		Id:          "iwu",
		Title:       "Nova Lista",
		Description: "So you can print the current date and time in a format that’s closer to what you may want to display to a user",
		Chooser:     chooser,
		Movies:      []Movie{movie},
		CreatedAt:   time.Now(),
		EditedAt:    time.Now(),
	}

	fmt.Println(list.Chooser)
}
