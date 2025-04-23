package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
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

func (u *GetListByUserIDUseCase) Execute(input GetListByUserIDInputDTO) (GetListByUserIDOutputDTO, []exceptions.ProblemDetails) {
	list, errGetList := u.ListRepository.GetListByID(input.ListID)
	if errGetList != nil {
		return GetListByUserIDOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching list",
				Status:   500,
				Detail:   "An error occurred while retrieving the list from the database.",
				Instance: exceptions.RFC500,
			},
		}
	}

	votes, errGetVotesByUserIDAndListID := u.VoteRepository.GetVotesByUserIDAndListID(input.UserID, input.ListID)
	if errGetVotesByUserIDAndListID != nil {
		return GetListByUserIDOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching votes",
				Status:   500,
				Detail:   "An error occurred while retrieving the votes for this user and list.",
				Instance: exceptions.RFC500,
			},
		}
	}

	numberOfVotes, errGetNumberOfVotesByListID := u.VoteRepository.GetNumberOfVotesByListID(input.ListID)
	if errGetNumberOfVotesByListID != nil {
		return GetListByUserIDOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching number of votes",
				Status:   500,
				Detail:   "An error occurred while retrieving the total number of votes for this list.",
				Instance: exceptions.RFC500,
			},
		}
	}

	rankItems, errGetRankItemsByVotes := u.VoteRepository.RankItemsByVotes(input.ListID, list.ListType)
	if errGetRankItemsByVotes != nil {
		return GetListByUserIDOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching ranked items",
				Status:   500,
				Detail:   "An error occurred while retrieving the ranked items for this list.",
				Instance: exceptions.RFC500,
			},
		}
	}

	outputRanking, err := list.FormatRanking(rankItems)
	if err != nil {
		return GetListByUserIDOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Invalid Input",
				Title:    "Invalid list type",
				Status:   400,
				Detail:   "The list type provided is invalid or cannot be processed.",
				Instance: exceptions.RFC400,
			},
		}
	}

	combinationsAlreadyVoted, errGetCombinationsAlreadyVoted := u.CombinationRepository.GetCombinationsAlreadyVoted(input.ListID)
	if errGetCombinationsAlreadyVoted != nil {
		return GetListByUserIDOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching combinations already voted",
				Status:   500,
				Detail:   "An error occurred while retrieving the combinations that have already been voted on.",
				Instance: exceptions.RFC500,
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
