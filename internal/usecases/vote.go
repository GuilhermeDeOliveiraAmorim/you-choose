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
	list, errGetListByID := u.ListRepository.GetListByID(input.Vote.ListID)
	if errGetListByID != nil {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching list",
				Detail:   "An error occurred while fetching the list from the database.",
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
				Detail:   "This vote has already been registered for the selected combination.",
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
					Detail:   "An error occurred while retrieving the movie information.",
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
					Detail:   "An error occurred while updating the movie's vote count.",
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
					Detail:   "An error occurred while retrieving the brand information.",
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
					Detail:   "An error occurred while updating the brand's vote count.",
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
				Detail:   "An error occurred while creating the vote entry in the database.",
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
