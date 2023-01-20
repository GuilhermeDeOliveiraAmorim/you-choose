package database

import (
	"database/sql"

	chooserEntity "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/domain/chooser/entity"
)

type ChooserRepository struct {
	Db *sql.DB
}

func NewChooserRepository(db *sql.DB) *ChooserRepository {
	return &ChooserRepository{
		Db: db,
	}
}

func (c *ChooserRepository) Save(chooser *chooserEntity.Chooser) error {
	stmt, err := c.Db.Prepare("INSERT INTO chooser (id, first_name, last_name, username, password, picture, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(chooser.ID, chooser.FirstName, chooser.LastName, chooser.UserName, chooser.Password, chooser.Picture, chooser.IsDeleted, chooser.CreatedAt, chooser.UpdatedAt, chooser.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}
