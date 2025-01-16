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
				Status:   404,
				Detail:   err.Error(),
				Instance: util.RFC404,
			},
		}
	} else if !user.Active {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    "User is not active",
				Status:   403,
				Detail:   "User is not active",
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
				Status:   500,
				Detail:   errGetListByID.Error(),
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
				Status:   409,
				Detail:   "Vote already registered for this combination.",
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
					Status:   500,
					Detail:   errGetMovie.Error(),
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
					Status:   500,
					Detail:   errUpdateWinner.Error(),
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
					Status:   500,
					Detail:   errGetBrand.Error(),
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
					Status:   500,
					Detail:   errUpdateWinner.Error(),
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
				Status:   500,
				Detail:   errVote.Error(),
				Instance: util.RFC500,
			},
		}
	}

	return VoteOutputDTO{
		SuccessMessage: "Vote created successfully!",
		ContentMessage: input.Vote.ListID,
	}, nil
}
