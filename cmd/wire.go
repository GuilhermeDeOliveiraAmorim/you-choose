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
