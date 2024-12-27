package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type GetListByIDInputDTO struct {
	ListID string `json:"list_id"`
	UserID string `json:"user_id"`
}

type GetListByIDOutputDTO struct {
	List  entities.List   `json:"list"`
	Votes []entities.Vote `json:"votes"`
}

type GetListByIDUseCase struct {
	ListRepository        repositories.ListRepository
	VoteRepository        repositories.VoteRepository
	CombinationRepository repositories.CombinationRepository
}

func NewGetListByIDUseCase(
	ListRepository repositories.ListRepository,
	VoteRepository repositories.VoteRepository,
	CombinationRepository repositories.CombinationRepository,
) *GetListByIDUseCase {
	return &GetListByIDUseCase{
		ListRepository:        ListRepository,
		VoteRepository:        VoteRepository,
		CombinationRepository: CombinationRepository,
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

	votes, errGetVotesByUserIDAndListID := u.VoteRepository.GetVotesByUserIDAndListID(input.UserID, input.ListID)
	if errGetVotesByUserIDAndListID != nil {
		return GetListByIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching votes",
				Status:   500,
				Detail:   errGetVotesByUserIDAndListID.Error(),
				Instance: util.RFC500,
			},
		}
	}

	return GetListByIDOutputDTO{
		List:  list,
		Votes: votes,
	}, nil
}
