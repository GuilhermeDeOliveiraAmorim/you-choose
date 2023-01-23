package database

import (
	"database/sql"
	"fmt"

	entity "github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type ChooserRepository struct {
	Db *sql.DB
}

func NewChooserRepository(db *sql.DB) *ChooserRepository {
	return &ChooserRepository{
		Db: db,
	}
}

func (c *ChooserRepository) Create(chooser *entity.Chooser) error {
	stmt, err := c.Db.Prepare("INSERT INTO choosers (id, first_name, last_name, username, password, picture, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)")

	fmt.Println("stmt")

	if err != nil {
		fmt.Print("o erro:", err)
		return err
	}
	_, err = stmt.Exec(chooser.ID, chooser.FirstName, chooser.LastName, chooser.UserName, chooser.Password, chooser.Picture, chooser.IsDeleted, chooser.CreatedAt, chooser.UpdatedAt, chooser.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}
