package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type VoteInputDTO struct {
	UserID        string `json:"user_id"`
	ListID        string `json:"list_id"`
	CombinationID string `json:"combination"`
	WinnerID      string `json:"winner_id"`
}

type VoteOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
}

type VoteUseCase struct {
	VoteRepository  repositories.VoteRepository
	ListRepository  repositories.ListRepository
	MovieRepository repositories.MovieRepository
}

func NewVoteUseCase(
	VoteRepository repositories.VoteRepository,
	ListRepository repositories.ListRepository,
	MovieRepository repositories.MovieRepository,
) *VoteUseCase {
	return &VoteUseCase{
		VoteRepository:  VoteRepository,
		ListRepository:  ListRepository,
		MovieRepository: MovieRepository,
	}
}

func (u *VoteUseCase) Execute(input VoteInputDTO) (VoteOutputDTO, []util.ProblemDetails) {
	listExists, errGetList := u.ListRepository.ThisListExistByID(input.ListID)
	if errGetList != nil {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching list",
				Status:   500,
				Detail:   errGetList.Error(),
				Instance: util.RFC500,
			},
		}
	}

	if !listExists {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Not Found",
				Status:   404,
				Detail:   "List not found.",
				Instance: util.RFC404,
			},
		}
	}

	_, errGetWinner := u.MovieRepository.GetMovieByID(input.WinnerID)
	if errGetWinner != nil {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching winner movie",
				Status:   500,
				Detail:   errGetWinner.Error(),
				Instance: util.RFC500,
			},
		}
	}

	voteAlreadyRegistered, errVoteAlreadyRegistered := u.VoteRepository.VoteAlreadyRegistered(input.UserID, input.CombinationID)
	if errVoteAlreadyRegistered != nil {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Conflict",
				Status:   409,
				Detail:   "Vote already registered for this combination.",
				Instance: util.RFC409,
			},
		}
	} else if voteAlreadyRegistered {
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

	newVote, newVoteErr := entities.NewVote(input.UserID, input.CombinationID, input.WinnerID)
	if newVoteErr != nil {
		return VoteOutputDTO{}, newVoteErr
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
		ContentMessage: input.ListID,
	}, nil
}
