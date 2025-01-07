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
	UserRepository        repositories.UserRepository
}

func NewGetListByIDUseCase(
	ListRepository repositories.ListRepository,
	VoteRepository repositories.VoteRepository,
	CombinationRepository repositories.CombinationRepository,
	UserRepository repositories.UserRepository,
) *GetListByIDUseCase {
	return &GetListByIDUseCase{
		ListRepository:        ListRepository,
		VoteRepository:        VoteRepository,
		CombinationRepository: CombinationRepository,
		UserRepository:        UserRepository,
	}
}

func (u *GetListByIDUseCase) Execute(input GetListByIDInputDTO) (GetListByIDOutputDTO, []util.ProblemDetails) {
	user, err := u.UserRepository.GetUser(input.UserID)
	if err != nil {
		return GetListByIDOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Not Found",
				Title:    "User not found",
				Status:   404,
				Detail:   err.Error(),
				Instance: util.RFC404,
			},
		}
	} else if !user.Active {
		return GetListByIDOutputDTO{}, []util.ProblemDetails{
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
