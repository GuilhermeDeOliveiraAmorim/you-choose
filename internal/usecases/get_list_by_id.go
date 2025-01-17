package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
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

func (u *GetListByIDUseCase) Execute(input GetListByIDInputDTO) (GetListByIDOutputDTO, []util.ProblemDetails) {
	list, errGetList := u.ListRepository.GetListByID(input.ListID)
	if errGetList != nil {
		return GetListByIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching list",
				Status:   500,
				Detail:   errGetList.Error(),
				Instance: util.RFC500,
			},
		}
	}

	numberOfVotes, errGetNumberOfVotesByListID := u.VoteRepository.GetNumberOfVotesByListID(input.ListID)
	if errGetNumberOfVotesByListID != nil {
		return GetListByIDOutputDTO{}, []util.ProblemDetails{
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
		return GetListByIDOutputDTO{}, []util.ProblemDetails{
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
		return GetListByIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Invalid Input",
				Title:    "Invalid list type",
				Status:   400,
				Detail:   err.Error(),
				Instance: util.RFC400,
			},
		}
	}

	return GetListByIDOutputDTO{
		List:          list,
		Ranking:       outputRanking,
		NumberOfVotes: numberOfVotes,
	}, nil
}
