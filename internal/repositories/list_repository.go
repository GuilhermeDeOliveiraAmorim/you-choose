package repositories

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"

type ListRepository interface {
	CreateList(list entities.List) error
	GetListByID(listID string) (entities.List, error)
	ThisListExistByName(listName string) (bool, error)
	ThisListExistByID(listID string) (bool, error)
}
