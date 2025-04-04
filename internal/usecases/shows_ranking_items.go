package usecases

import (
	"sort"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type ShowsRankingItemsInputDTO struct {
	ListType string `json:"list_type"`
}

type ShowsRankingItemsOutputDTO struct {
	Ranking []interface{} `json:"ranking"`
}

type ShowsRankingItemsUseCase struct {
	MovieRepository repositories.MovieRepository
	BrandRepository repositories.BrandRepository
}

func NewShowsRankingItemsUseCase(
	MovieRepository repositories.MovieRepository,
	BrandRepository repositories.BrandRepository,
) *ShowsRankingItemsUseCase {
	return &ShowsRankingItemsUseCase{
		MovieRepository: MovieRepository,
		BrandRepository: BrandRepository,
	}
}

func (u *ShowsRankingItemsUseCase) Execute(input ShowsRankingItemsInputDTO) (ShowsRankingItemsOutputDTO, []util.ProblemDetails) {
	var ranking []interface{}

	switch input.ListType {
	case entities.MOVIE_TYPE:
		movies, err := u.MovieRepository.GetMovies()
		if err != nil {
			return ShowsRankingItemsOutputDTO{}, []util.ProblemDetails{
				{
					Type:     "Internal Server Error",
					Title:    "Error fetching movies",
					Detail:   err.Error(),
					Status:   500,
					Instance: util.RFC500,
				},
			}
		}

		sort.Slice(movies, func(i, j int) bool {
			return movies[i].VotesCount > movies[j].VotesCount
		})

		ranking = make([]interface{}, len(movies))

		for i, movie := range movies {
            ranking[i] = movie
        }

	case entities.BRAND_TYPE:
		brands, err := u.BrandRepository.GetBrands()
		if err != nil {
			return ShowsRankingItemsOutputDTO{}, []util.ProblemDetails{
				{
					Type:     "Internal Server Error",
					Title:    "Error fetching brands",
					Detail:   err.Error(),
					Status:   500,
					Instance: util.RFC500,
				},
			}
		}

		sort.Slice(brands, func(i, j int) bool {
			return brands[i].VotesCount > brands[j].VotesCount
		})

		ranking = make([]interface{}, len(brands))

		for i, brand := range brands {
            ranking[i] = brand
        }
	}

	return ShowsRankingItemsOutputDTO{
		Ranking: ranking,
	}, nil
}
