package createmovielist

import (
	movieList "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/movie-list/entity"
)

func CreateMovieList(input *InputCreateMovieListDto) *OutputCreateMovieListDto {
	if input.Title == "" {
		return nil
	}

	if input.Description == "" {
		return nil
	}

	if input.Picture == "" {
		return nil
	}

	movieListOutput, _ := movieList.NewMovieList(input.Title, input.Description, input.Picture)

	output := OutputCreateMovieListDto{
		movieListOutput.ID,
		movieListOutput.Title,
		movieListOutput.Description,
		movieListOutput.Picture,
	}

	return &output
}
