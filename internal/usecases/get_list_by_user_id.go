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
	List                entities.List          `json:"list"`
	NumberOfVotes       int                    `json:"number_of_votes"`
	Ranking             interface{}            `json:"ranking"`
	VotedCombinations   []entities.Combination `json:"voted_combinations"`
	UnvotedCombinations []entities.Combination `json:"unvoted_combinations"`
	Votes               []entities.Vote        `json:"votes"`
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

	rankItems, errGetRankItemsByVotes := u.VoteRepository.RankItemsByVotes(input.ListID, list.ListType)
	if errGetRankItemsByVotes != nil {
		return GetListByUserIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching ranked items",
				Status:   500,
				Detail:   errGetRankItemsByVotes.Error(),
				Instance: util.RFC500,
			},
		}
	}

	outputRanking, err := list.FormatRanking(rankItems)
	if err != nil {
		return GetListByUserIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Invalid Input",
				Title:    "Invalid list type",
				Status:   400,
				Detail:   err.Error(),
				Instance: util.RFC400,
			},
		}
	}

	combinationsAlreadyVoted, errGetCombinationsAlreadyVoted := u.CombinationRepository.GetCombinationsAlreadyVoted(input.ListID)
	if errGetCombinationsAlreadyVoted != nil {
		return GetListByUserIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching combinations already voted",
				Status:   500,
				Detail:   errGetCombinationsAlreadyVoted.Error(),
				Instance: util.RFC500,
			},
		}
	}

	var unvotedCombinations []entities.Combination

	for _, combination := range list.Combinations {
		isVoted := false
		for _, votedCombination := range combinationsAlreadyVoted {
			if combination.ID == votedCombination.ID {
				isVoted = true
				break
			}
		}

		if !isVoted {
			unvotedCombinations = append(unvotedCombinations, combination)
		}
	}

	return GetListByUserIDOutputDTO{
		List:                list,
		NumberOfVotes:       numberOfVotes,
		Ranking:             outputRanking,
		VotedCombinations:   combinationsAlreadyVoted,
		UnvotedCombinations: unvotedCombinations,
		Votes:               votes,
	}, nil
}
