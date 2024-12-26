package gorm_implementation

import (
	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entities"
	"gorm.io/gorm"
)

type ListRepository struct {
	gorm *gorm.DB
}

func NewListRepository(gorm *gorm.DB) *ListRepository {
	return &ListRepository{
		gorm: gorm,
	}
}

func (c *ListRepository) CreateList(list entities.List) error {
	tx := c.gorm.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := tx.Create(&Lists{
		ID:            list.ID,
		Active:        list.Active,
		CreatedAt:     list.CreatedAt,
		UpdatedAt:     list.UpdatedAt,
		DeactivatedAt: list.DeactivatedAt,
		Name:          list.Name,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, movie := range list.Movies {
		if err := tx.Exec("INSERT INTO list_tags (lists_id, tags_id) VALUES (?, ?)", list.ID, movie.ID).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
