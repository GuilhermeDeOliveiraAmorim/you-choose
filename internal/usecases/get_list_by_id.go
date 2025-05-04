package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/exceptions"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
)

type GetListByIDInputDTO struct {
	ListID string `json:"list_id"`
}

type GetListByIDOutputDTO struct {
	List          entities.List `json:"list"`
	Ranking       interface{}   `json:"ranking"`
	NumberOfVotes int           `json:"number_of_votes"`
}

type GetListByIDUseCase struct {
	ListRepository repositories.ListRepository
	VoteRepository repositories.VoteRepository
}

func NewGetListByIDUseCase(
	ListRepository repositories.ListRepository,
	VoteRepository repositories.VoteRepository,
) *GetListByIDUseCase {
	return &GetListByIDUseCase{
		ListRepository: ListRepository,
		VoteRepository: VoteRepository,
	}
}

func (u *GetListByIDUseCase) Execute(input GetListByIDInputDTO) (GetListByIDOutputDTO, []exceptions.ProblemDetails) {
	list, errGetList := u.ListRepository.GetListByID(input.ListID)
	if errGetList != nil {
		if errGetList.Error() == "list not found" {
			return GetListByIDOutputDTO{}, []exceptions.ProblemDetails{
				{
					Type:     "Not Found",
					Title:    "List not found",
					Status:   404,
					Detail:   "The requested list was not found.",
					Instance: exceptions.RFC404,
				},
			}
		}

		return GetListByIDOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching list",
				Status:   500,
				Detail:   "An error occurred while retrieving the list from the database.",
				Instance: exceptions.RFC500,
			},
		}
	}

	numberOfVotes, errGetNumberOfVotesByListID := u.VoteRepository.GetNumberOfVotesByListID(input.ListID)
	if errGetNumberOfVotesByListID != nil {
		return GetListByIDOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching number of votes",
				Status:   500,
				Detail:   "An error occurred while retrieving the total number of votes for the list.",
				Instance: exceptions.RFC500,
			},
		}
	}

	rankItems, errGetRankItemsByVotes := u.VoteRepository.RankItemsByVotes(input.ListID, list.ListType)
	if errGetRankItemsByVotes != nil {
		return GetListByIDOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching ranked items",
				Status:   500,
				Detail:   "An error occurred while retrieving the ranked items for the list.",
				Instance: exceptions.RFC500,
			},
		}
	}

	outputRanking, err := list.FormatRanking(rankItems)
	if err != nil {
		return GetListByIDOutputDTO{}, []exceptions.ProblemDetails{
			{
				Type:     "Invalid Input",
				Title:    "Invalid list type",
				Status:   400,
				Detail:   "The list type is invalid or cannot be processed.",
				Instance: exceptions.RFC400,
			},
		}
	}

	return GetListByIDOutputDTO{
		List:          list,
		Ranking:       outputRanking,
		NumberOfVotes: numberOfVotes,
	}, nil
}
