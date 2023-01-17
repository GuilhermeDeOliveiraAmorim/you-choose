package deletechooser

import (
	"errors"
	"time"

	chooserRepository "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/repository"
)

type DeleteChooserUseCase struct {
	ChooserRepository chooserRepository.ChooserRepositoryInterface
}

func NewDeleteChooserUseCase(ChooserRepository chooserRepository.ChooserRepositoryInterface) *DeleteChooserUseCase {
	return &DeleteChooserUseCase{
		ChooserRepository: ChooserRepository,
	}
}

func (chooserUseCase *DeleteChooserUseCase) Execute(input InputDeleteChooserDto) (OutputDeleteChooserDto, error) {
	chooser, err := chooserUseCase.ChooserRepository.FindById(input.ID)

	output := OutputDeleteChooserDto{}

	if err != nil {
		return output, errors.New("chooser not found")
	}

	chooser.IsDeleted = true
	chooser.DeletedAt = time.Now()

	output.IsDeleted = chooser.IsDeleted

	return output, err
}
