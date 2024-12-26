package repositories

import "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"

type ListRepository interface {
	CreateList(list entities.List) error
	GetListByID(listID string) (entities.List, error)
	GetLists() ([]entities.List, error)
}
