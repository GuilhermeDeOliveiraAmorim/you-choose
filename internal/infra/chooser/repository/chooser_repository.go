package repository

import (
	chooser "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/entity"
	chooserRepository "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/repository"
)

type ChooserRepository interface {
	chooserRepository.ChooserRepositoryInterface
}

func Add(chooser *chooser.Chooser, c ChooserRepository) {

}

func FindById(id string, c ChooserRepository) {

}
