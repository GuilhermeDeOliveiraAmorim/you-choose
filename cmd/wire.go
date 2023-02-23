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

var SetMovieRepositoryDependency = wire.NewSet(
	database.NewMovieRepository,
	wire.Bind(new(*entity.MovieRepositoryInterface), new(*database.MovieRepository)),
)

var SetActorRepositoryDependency = wire.NewSet(
	database.NewActorRepository,
	wire.Bind(new(*entity.ActorRepositoryInterface), new(*database.ActorRepository)),
)

var SetWriterRepositoryDependency = wire.NewSet(
	database.NewWriterRepository,
	wire.Bind(new(*entity.WriterRepositoryInterface), new(*database.WriterRepository)),
)

var SetDirectorRepositoryDependency = wire.NewSet(
	database.NewDirectorRepository,
	wire.Bind(new(*entity.DirectorRepositoryInterface), new(*database.DirectorRepository)),
)

var SetGenreRepositoryDependency = wire.NewSet(
	database.NewGenreRepository,
	wire.Bind(new(*entity.GenreRepositoryInterface), new(*database.GenreRepository)),
)

var SetTagRepositoryDependency = wire.NewSet(
	database.NewTagRepository,
	wire.Bind(new(*entity.TagRepositoryInterface), new(*database.TagRepository)),
)

var SetFileRepositoryDependency = wire.NewSet(
	database.NewFileRepository,
	wire.Bind(new(*entity.FileRepositoryInterface), new(*database.FileRepository)),
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

func NewCreateMovieUseCase(db *sql.DB) *usecases.MovieUseCase {
	wire.Build(
		SetMovieRepositoryDependency,
	)
	return &usecases.MovieUseCase{}
}

func NewWebMovieHandler(db *sql.DB) *web.WebMovieHandler {
	wire.Build(
		SetMovieRepositoryDependency,
	)
	return &web.WebMovieHandler{}
}

func NewCreateActorUseCase(db *sql.DB) *usecases.ActorUseCase {
	wire.Build(
		SetActorRepositoryDependency,
	)
	return &usecases.ActorUseCase{}
}

func NewWebActorHandler(db *sql.DB) *web.WebActorHandler {
	wire.Build(
		SetActorRepositoryDependency,
	)
	return &web.WebActorHandler{}
}

func NewCreateWriterUseCase(db *sql.DB) *usecases.WriterUseCase {
	wire.Build(
		SetWriterRepositoryDependency,
	)
	return &usecases.WriterUseCase{}
}

func NewWebWriterHandler(db *sql.DB) *web.WebWriterHandler {
	wire.Build(
		SetWriterRepositoryDependency,
	)
	return &web.WebWriterHandler{}
}

func NewCreateDirectorUseCase(db *sql.DB) *usecases.DirectorUseCase {
	wire.Build(
		SetDirectorRepositoryDependency,
	)
	return &usecases.DirectorUseCase{}
}

func NewWebDirectorHandler(db *sql.DB) *web.WebDirectorHandler {
	wire.Build(
		SetDirectorRepositoryDependency,
	)
	return &web.WebDirectorHandler{}
}

func NewCreateGenreUseCase(db *sql.DB) *usecases.GenreUseCase {
	wire.Build(
		SetGenreRepositoryDependency,
	)
	return &usecases.GenreUseCase{}
}

func NewWebGenreHandler(db *sql.DB) *web.WebGenreHandler {
	wire.Build(
		SetGenreRepositoryDependency,
	)
	return &web.WebGenreHandler{}
}

func NewCreateTagUseCase(db *sql.DB) *usecases.TagUseCase {
	wire.Build(
		SetTagRepositoryDependency,
	)
	return &usecases.TagUseCase{}
}

func NewWebTagHandler(db *sql.DB) *web.WebTagHandler {
	wire.Build(
		SetTagRepositoryDependency,
	)
	return &web.WebTagHandler{}
}

func NewCreateFileUseCase(db *sql.DB) *usecases.FileUseCase {
	wire.Build(
		SetFileRepositoryDependency,
	)
	return &usecases.FileUseCase{}
}

func NewWebFileHandler(db *sql.DB) *web.WebFileHandler {
	wire.Build(
		SetFileRepositoryDependency,
	)
	return &web.WebFileHandler{}
}
