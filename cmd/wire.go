package main

import (
	"database/sql"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infra/web"
	webserver "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infra/web"

	dbChooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/infra/database"

	chooserCreateUseCase "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/application/usecases/chooser/create_chooser"

	chooserRepository "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/repository"

	"github.com/google/wire"
)

var SetChooserRepositoryDependency = wire.NewSet(
	dbChooser.NewChooserRepository,
	wire.Bind(new(*chooserRepository.ChooserRepositoryInterface), new(*dbChooser.ChooserRepository)),
)

func NewCreateChooserUseCase(db *sql.DB) *chooserCreateUseCase.CreateChooserUseCase {
	wire.Build(
		SetChooserRepositoryDependency,
	)
	return &chooserCreateUseCase.CreateChooserUseCase{}
}

func NewWebChooserHandler(db *sql.DB) *web.WebChooserHandler {
	wire.Build(
		SetChooserRepositoryDependency,
	)
	return &webserver.WebChooserHandler{}
}
