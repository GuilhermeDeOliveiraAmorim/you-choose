package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type VoteInputDTO struct {
	ListID        string `json:"list_id"`
	FirstMovieID  string `json:"first_movie_id"`
	SecondMovieID string `json:"second_movie_id"`
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

	_, errGetFirstMovie := u.MovieRepository.GetMovieByID(input.FirstMovieID)
	if errGetFirstMovie != nil {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching first movie",
				Status:   500,
				Detail:   errGetFirstMovie.Error(),
				Instance: util.RFC500,
			},
		}
	}

	_, errGetSecondMovie := u.MovieRepository.GetMovieByID(input.SecondMovieID)
	if errGetSecondMovie != nil {
		return VoteOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching second movie",
				Status:   500,
				Detail:   errGetSecondMovie.Error(),
				Instance: util.RFC500,
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

	newVote, newVoteErr := entities.NewVote(input.ListID, input.FirstMovieID, input.SecondMovieID, input.WinnerID)
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
