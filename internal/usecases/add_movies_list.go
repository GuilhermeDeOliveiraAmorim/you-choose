package usecases

import (
	"fmt"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type AddMoviesListInputDTO struct {
	ListID string   `json:"name"`
	Movies []string `json:"movies"`
}

type AddMoviesListOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
}

type AddMoviesListUseCase struct {
	ListRepository  repositories.ListRepository
	MovieRepository repositories.MovieRepository
}

func NewAddMoviesListUseCase(
	ListRepository repositories.ListRepository,
	MovieRepository repositories.MovieRepository,
) *AddMoviesListUseCase {
	return &AddMoviesListUseCase{
		ListRepository:  ListRepository,
		MovieRepository: MovieRepository,
	}
}

func (u *AddMoviesListUseCase) Execute(input AddMoviesListInputDTO) (AddMoviesListOutputDTO, []util.ProblemDetails) {
	var problems []util.ProblemDetails

	list, errGetList := u.ListRepository.GetListByID(input.ListID)
	if errGetList != nil {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching list",
				Status:   500,
				Detail:   errGetList.Error(),
				Instance: util.RFC500,
			},
		}
	}

	for _, movieID := range input.Movies {
		for _, movie := range list.Movies {
			if movie.ID == movieID {
				problems = append(problems,
					util.ProblemDetails{
						Type:     "Validation Error",
						Title:    "Movie already in list",
						Status:   400,
						Detail:   fmt.Sprintf("Movie with ID %s already exists in the list.", movieID),
						Instance: util.RFC400,
					},
				)
			}
		}
	}

	if len(problems) > 0 {
		return AddMoviesListOutputDTO{}, problems
	}

	movies, errGetMoviesByID := u.MovieRepository.GetMoviesByIDs(input.Movies)
	if errGetMoviesByID != nil {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching movies",
				Status:   500,
				Detail:   errGetMoviesByID.Error(),
				Instance: util.RFC500,
			},
		}
	}

	list.AddMovies(movies)

	errAddMovies := u.ListRepository.AddMovies(list)
	if errAddMovies != nil {
		return AddMoviesListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error adding movies to list",
				Status:   500,
				Detail:   errAddMovies.Error(),
				Instance: util.RFC500,
			},
		}
	}

	return AddMoviesListOutputDTO{
		SuccessMessage: "Movies added successfully.",
		ContentMessage: "The movies were successfully added to the list.",
	}, nil
}
