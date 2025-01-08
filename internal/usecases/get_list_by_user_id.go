package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type GetListByUserIDInputDTO struct {
	ListID string `json:"list_id"`
	UserID string `json:"user_id"`
}

type GetListByUserIDOutputDTO struct {
	List          entities.List    `json:"list"`
	NumberOfVotes int              `json:"number_of_votes"`
	RankMovies    []entities.Movie `json:"rank_movies"`
	Votes         []entities.Vote  `json:"votes"`
}

type GetListByUserIDUseCase struct {
	ListRepository        repositories.ListRepository
	VoteRepository        repositories.VoteRepository
	CombinationRepository repositories.CombinationRepository
	UserRepository        repositories.UserRepository
}

func NewGetListByUserIDUseCase(
	ListRepository repositories.ListRepository,
	VoteRepository repositories.VoteRepository,
	CombinationRepository repositories.CombinationRepository,
	UserRepository repositories.UserRepository,
) *GetListByUserIDUseCase {
	return &GetListByUserIDUseCase{
		ListRepository:        ListRepository,
		VoteRepository:        VoteRepository,
		CombinationRepository: CombinationRepository,
		UserRepository:        UserRepository,
	}
}

func (u *GetListByUserIDUseCase) Execute(input GetListByUserIDInputDTO) (GetListByUserIDOutputDTO, []util.ProblemDetails) {
	user, err := u.UserRepository.GetUser(input.UserID)
	if err != nil {
		return GetListByUserIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Not Found",
				Title:    "User not found",
				Status:   404,
				Detail:   err.Error(),
				Instance: util.RFC404,
			},
		}
	} else if !user.Active {
		return GetListByUserIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Forbidden",
				Title:    "User is not active",
				Status:   403,
				Detail:   "User is not active",
				Instance: util.RFC403,
			},
		}
	}

	list, errGetList := u.ListRepository.GetListByID(input.ListID)
	if errGetList != nil {
		return GetListByUserIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching list",
				Status:   500,
				Detail:   errGetList.Error(),
				Instance: util.RFC500,
			},
		}
	}

	votes, errGetVotesByUserIDAndListID := u.VoteRepository.GetVotesByUserIDAndListID(input.UserID, input.ListID)
	if errGetVotesByUserIDAndListID != nil {
		return GetListByUserIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching votes",
				Status:   500,
				Detail:   errGetVotesByUserIDAndListID.Error(),
				Instance: util.RFC500,
			},
		}
	}

	numberOfVotes, errGetNumberOfVotesByListID := u.VoteRepository.GetNumberOfVotesByListID(input.ListID)
	if errGetNumberOfVotesByListID != nil {
		return GetListByUserIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching number of votes",
				Status:   500,
				Detail:   errGetNumberOfVotesByListID.Error(),
				Instance: util.RFC500,
			},
		}
	}

	rankMovies, errGetRankMoviesByVotes := u.VoteRepository.RankMoviesByVotes(input.ListID)
	if errGetRankMoviesByVotes != nil {
		return GetListByUserIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching ranked movies",
				Status:   500,
				Detail:   errGetRankMoviesByVotes.Error(),
				Instance: util.RFC500,
			},
		}
	}

	return GetListByUserIDOutputDTO{
		List:          list,
		NumberOfVotes: numberOfVotes,
		RankMovies:    rankMovies,
		Votes:         votes,
	}, nil
}
