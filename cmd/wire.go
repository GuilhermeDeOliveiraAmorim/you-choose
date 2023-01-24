package main

import (
	"database/sql"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infra/web"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infra/database"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/usecases"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"

	"github.com/google/wire"
)

var SetChooserRepositoryDependency = wire.NewSet(
	database.NewChooserRepository,
	wire.Bind(new(*entity.ChooserRepositoryInterface), new(*database.ChooserRepository)),
)

var SetMovieListRepositoryDependency = wire.NewSet(
	database.NewMovieListRepository,
	wire.Bind(new(*entity.MovieListRepositoryInterface), new(*database.MovieListRepository)),
)

func NewCreateChooserUseCase(db *sql.DB) *usecases.ChooserUseCase {
	wire.Build(
		SetChooserRepositoryDependency,
	)
	return &usecases.ChooserUseCase{}
}

func NewWebChooserHandler(db *sql.DB) *web.WebChooserHandler {
	wire.Build(
		SetChooserRepositoryDependency,
	)
	return &web.WebChooserHandler{}
}

func NewCreateMovieListUseCase(db *sql.DB) *usecases.MovieListUseCase {
	wire.Build(
		SetMovieListRepositoryDependency,
	)
	return &usecases.MovieListUseCase{}
}

func NewWebMovieListHandler(db *sql.DB) *web.WebMovieListHandler {
	wire.Build(
		SetMovieListRepositoryDependency,
	)
	return &web.WebMovieListHandler{}
}
