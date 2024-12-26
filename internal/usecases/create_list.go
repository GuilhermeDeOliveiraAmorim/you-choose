package usecases

import (
	"strings"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/repositories"
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/util"
)

type CreateListInputDTO struct {
	Name   string           `json:"name"`
	Movies []entities.Movie `json:"movies"`
}

type CreateListOutputDTO struct {
	SuccessMessage string `json:"success_message"`
	ContentMessage string `json:"content_message"`
}

type CreateListUseCase struct {
	ListRepository repositories.ListRepository
}

func NewCreateListUseCase(
	ListRepository repositories.ListRepository,
) *CreateListUseCase {
	return &CreateListUseCase{
		ListRepository: ListRepository,
	}
}

func (u *CreateListUseCase) Execute(input CreateListInputDTO) (CreateListOutputDTO, []util.ProblemDetails) {
	listExists, errThisListExist := u.ListRepository.ThisListExistByName(input.Name)
	if errThisListExist != nil && strings.Compare(errThisListExist.Error(), "list not found") > 0 {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error fetching existing list",
				Status:   500,
				Detail:   errThisListExist.Error(),
				Instance: util.RFC500,
			},
		}
	}

	if listExists {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Conflict",
				Status:   409,
				Detail:   "A list with the same name already exists.",
				Instance: util.RFC409,
			},
		}
	}

	list, problems := entities.NewList(input.Name)
	if len(problems) > 0 {
		return CreateListOutputDTO{}, problems
	}

	if len(input.Movies) > 0 {
		list.AddMovies(input.Movies)
	} else {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Validation Error",
				Title:    "Bad Request",
				Status:   400,
				Detail:   "At least one movie must be provided.",
				Instance: util.RFC400,
			},
		}
	}

	errCreateList := u.ListRepository.CreateList(*list)
	if errCreateList != nil {
		return CreateListOutputDTO{}, []util.ProblemDetails{
			{
				Type:     "Internal Server Error",
				Title:    "Error creating list",
				Status:   500,
				Detail:   errCreateList.Error(),
				Instance: util.RFC500,
			},
		}
	}

	return CreateListOutputDTO{
		SuccessMessage: "List created successfully!",
		ContentMessage: list.Name,
	}, nil
}
