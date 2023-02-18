package database

import (
	"database/sql"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type TagRepository struct {
	Db *sql.DB
}

func NewTagRepository(db *sql.DB) *TagRepository {
	return &TagRepository{
		Db: db,
	}
}

func (tagRepository *TagRepository) Create(tag *entity.Tag) error {
	stmt, err := tagRepository.Db.Prepare("INSERT INTO tags (tag_id, name, picture, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(tag.ID, tag.Name, tag.Picture, tag.IsDeleted, tag.CreatedAt, tag.UpdatedAt, tag.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (tagRepository *TagRepository) Find(id string) (entity.Tag, error) {
	var tag entity.Tag

	rows, err := tagRepository.Db.Query("SELECT * FROM tags WHERE tag_id = $1", id)
	if err != nil {
		return tag, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.Picture,
			&tag.IsDeleted,
			&tag.CreatedAt,
			&tag.UpdatedAt,
			&tag.DeletedAt); err != nil {
			return tag, err
		}
	}

	if err = rows.Err(); err != nil {
		return tag, err
	}

	return tag, nil
}

func (tagRepository *TagRepository) Update(tag *entity.Tag) error {
	stmt, err := tagRepository.Db.Prepare("UPDATE tags SET name = $1, picture = $2, is_deleted = $3, created_at = $4, updated_at = $5, deleted_at = $6 WHERE id = $7")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(tag.Name, tag.Picture, tag.IsDeleted, tag.CreatedAt, tag.UpdatedAt, tag.DeletedAt, tag.ID)
	if err != nil {
		return err
	}

	return nil
}

func (tagRepository *TagRepository) Delete(tag *entity.Tag) error {
	stmt, err := tagRepository.Db.Prepare("UPDATE tags SET is_deleted = $1, deleted_at = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(tag.IsDeleted, tag.DeletedAt, tag.ID)
	if err != nil {
		return err
	}

	return nil
}

func (tagRepository *TagRepository) IsDeleted(id string) error {
	var tag entity.Tag

	rows, err := tagRepository.Db.Query("SELECT * FROM tags WHERE id = $1", id)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Picture, &tag.IsDeleted, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeletedAt); err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (tagRepository *TagRepository) FindAll() ([]entity.Tag, error) {
	rows, err := tagRepository.Db.Query("SELECT tag_id, name, picture, is_deleted, created_at, updated_at, deleted_at FROM tags")
	if err != nil {
		return nil, err
	}

	var tags []entity.Tag

	for rows.Next() {
		var tag entity.Tag

		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Picture, &tag.IsDeleted, &tag.CreatedAt, &tag.UpdatedAt, &tag.DeletedAt); err != nil {
			return tags, err
		}

		tags = append(tags, tag)
	}

	if err = rows.Err(); err != nil {
		return tags, err
	}

	return tags, nil
}

func (tagRepository *TagRepository) FindTagByName(name string) (entity.Tag, error) {
	var tag entity.Tag

	rows, err := tagRepository.Db.Query("SELECT * FROM tags WHERE name = $1", name)
	if err != nil {
		return tag, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.Picture,
			&tag.IsDeleted,
			&tag.CreatedAt,
			&tag.UpdatedAt,
			&tag.DeletedAt); err != nil {
			return tag, err
		}
	}

	if err = rows.Err(); err != nil {
		return tag, err
	}

	return tag, nil
}
