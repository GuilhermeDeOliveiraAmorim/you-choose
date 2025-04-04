package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type Vote struct {
	ListID        string `json:"list_id"`
	CombinationID string `json:"combination_id"`
	WinnerID      string `json:"winner_id"`
}

type VoteInputDTO struct {
	UserID string `json:"user_id"`
	Vote   Vote   `json:"vote"`
}

type VoteOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
}

type VoteUseCase struct {
	VoteRepository  repositories.VoteRepository
	ListRepository  repositories.ListRepository
	MovieRepository repositories.MovieRepository
	UserRepository  repositories.UserRepository
	BrandRepository repositories.BrandRepository
}

func NewVoteUseCase(
	VoteRepository repositories.VoteRepository,
	ListRepository repositories.ListRepository,
	MovieRepository repositories.MovieRepository,
	UserRepository repositories.UserRepository,
	BrandRepository repositories.BrandRepository,
) *VoteUseCase {
	return &VoteUseCase{
		VoteRepository:  VoteRepository,
		ListRepository:  ListRepository,
		MovieRepository: MovieRepository,
		UserRepository:  UserRepository,
		BrandRepository: BrandRepository,
	}
}

func (u *VoteUseCase) Execute(input VoteInputDTO) (VoteOutputDTO, []util.ProblemDetails) {
	user, err := u.UserRepository.GetUser(input.UserID)
	if err != nil {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Not Found",
				Title:    "User not found",
				Detail:   err.Error(),
				Status:   404,
				Instance: util.RFC404,
			},
		}
	} else if !user.Active {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    "User is not active",
				Detail:   "User is not active",
				Status:   403,
				Instance: util.RFC403,
			},
		}
	}

	list, errGetListByID := u.ListRepository.GetListByID(input.Vote.ListID)
	if errGetListByID != nil {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching list",
				Detail:   errGetListByID.Error(),
				Status:   500,
				Instance: util.RFC500,
			},
		}
	}

	voteAlreadyRegistered, errVoteAlreadyRegistered := u.VoteRepository.VoteAlreadyRegistered(input.UserID, input.Vote.CombinationID)
	if (errVoteAlreadyRegistered != nil) || voteAlreadyRegistered {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Conflict",
				Detail:   "Vote already registered for this combination",
				Status:   409,
				Instance: util.RFC409,
			},
		}
	}

	newVote, newVoteErr := entities.NewVote(input.UserID, input.Vote.CombinationID, input.Vote.WinnerID)
	if newVoteErr != nil {
		return VoteOutputDTO{}, newVoteErr
	}

	switch list.ListType {
	case entities.MOVIE_TYPE:
		movie, errGetMovie := u.MovieRepository.GetMovieByID(input.Vote.WinnerID)
		if errGetMovie != nil {
			return VoteOutputDTO{}, []util.ProblemDetails{
				{
					Type:     "Internal Server Error",
					Title:    "Error fetching winner movie",
					Detail:   errGetMovie.Error(),
					Status:   500,
					Instance: util.RFC500,
				},
			}
		}

		movie.IncrementVotesCount()

		errUpdateWinner := u.MovieRepository.UpdadeMovie(movie)
		if errUpdateWinner != nil {
			return VoteOutputDTO{}, []util.ProblemDetails{
				{
					Type:     "Internal Server Error",
					Title:    "Error updating winner movie",
					Detail:   errUpdateWinner.Error(),
					Status:   500,
					Instance: util.RFC500,
				},
			}
		}

	case entities.BRAND_TYPE:
		brand, errGetBrand := u.BrandRepository.GetBrandByID(input.Vote.WinnerID)
		if errGetBrand != nil {
			return VoteOutputDTO{}, []util.ProblemDetails{
				{
					Type:     "Internal Server Error",
					Title:    "Error fetching winner brand",
					Detail:   errGetBrand.Error(),
					Status:   500,
					Instance: util.RFC500,
				},
			}
		}

		brand.IncrementVotesCount()

		errUpdateWinner := u.BrandRepository.UpdadeBrand(brand)
		if errUpdateWinner != nil {
			return VoteOutputDTO{}, []util.ProblemDetails{
				{
					Type:     "Internal Server Error",
					Title:    "Error updating winner brand",
					Detail:   errUpdateWinner.Error(),
					Status:   500,
					Instance: util.RFC500,
				},
			}
		}
	}

	errVote := u.VoteRepository.CreateVote(*newVote)
	if errVote != nil {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error creating vote",
				Detail:   errVote.Error(),
				Status:   500,
				Instance: util.RFC500,
			},
		}
	}

	return VoteOutputDTO{
		SuccessMessage: "Vote created successfully!",
		ContentMessage: input.Vote.ListID,
	}, nil
}
