package usecases

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type SimpleList struct {
	entities.SharedEntity
	Name     string `json:"name"`
	Cover    string `json:"cover"`
	ListType string `json:"list_type"`
}

type GetListsInputDTO struct{}

type GetListsOutputDTO struct {
	Lists []SimpleList `json:"lists"`
}

type GetListsUseCase struct {
	ListRepository repositories.ListRepository
}

func NewGetListsUseCase(
	ListRepository repositories.ListRepository,
) *GetListsUseCase {
	return &GetListsUseCase{
		ListRepository: ListRepository,
	}
}

func (u *GetListsUseCase) Execute(input GetListsInputDTO) (GetListsOutputDTO, []util.ProblemDetails) {
	lists, errGetLists := u.ListRepository.GetLists()
	if errGetLists != nil {
		return GetListsOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching lists",
				Status:   500,
				Detail:   errGetLists.Error(),
				Instance: util.RFC500,
			},
		}
	}

	var simpleLists []SimpleList

	for _, list := range lists {
		simpleLists = append(simpleLists, SimpleList{
			SharedEntity: list.SharedEntity,
			Name:         list.Name,
			Cover:        list.Cover,
			ListType:     list.ListType,
		})
	}

	return GetListsOutputDTO{
		Lists: simpleLists,
	}, nil
}
