// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"

	// "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infra/database"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infra/web"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"
	// "github.com/google/wire"
)

func NewCreateChooserUseCaseGen(db *sql.DB) *usecases.ChooserUseCase {
	chooserRepository := database.NewChooserRepository(db)
	movieListRepository := database.NewMovieListRepository(db)
	chooserUseCase := usecases.NewChooserUseCase(chooserRepository, movieListRepository)
	return chooserUseCase
}

func NewWebChooserHandlerGen(db *sql.DB) *web.WebChooserHandler{
	chooserRepository := database.NewChooserRepository(db)
	movieListRepository := database.NewMovieListRepository(db)
	webChooserHandler := web.NewChooserHandler(chooserRepository, movieListRepository)
	return webChooserHandler
}

func NewCreateMovieListUseCaseGen(db *sql.DB) *usecases.MovieListUseCase {
	movieListRepository := database.NewMovieListRepository(db)
	chooserRepository := database.NewChooserRepository(db)
	MovieListUseCase := usecases.NewMovieListUseCase(movieListRepository, chooserRepository)
	return MovieListUseCase
}

func NewWebMovieListHandlerGen(db *sql.DB) *web.WebMovieListHandler{
	movieListRepository := database.NewMovieListRepository(db)
	chooserRepository := database.NewChooserRepository(db)
	webMovieListHandler := web.NewMovieListHandler(movieListRepository, chooserRepository)
	return webMovieListHandler
}

func NewCreateMovieUseCaseGen(db *sql.DB) *usecases.MovieUseCase {
	movieRepository := database.NewMovieRepository(db)
	MovieUseCase := usecases.NewMovieUseCase(movieRepository)
	return MovieUseCase
}

func NewWebMovieHandlerGen(db *sql.DB) *web.WebMovieHandler{
	movieRepository := database.NewMovieRepository(db)
	webMovieHandler := web.NewMovieHandler(movieRepository)
	return webMovieHandler
}