package database

import (
	"database/sql"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type WriterRepository struct {
	Db *sql.DB
}

func NewWriterRepository(db *sql.DB) *WriterRepository {
	return &WriterRepository{
		Db: db,
	}
}

func (writerRepository *WriterRepository) Create(writer *entity.Writer) error {
	stmt, err := writerRepository.Db.Prepare("INSERT INTO writers (writer_id, name, picture, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(writer.ID, writer.Name, writer.Picture, writer.IsDeleted, writer.CreatedAt, writer.UpdatedAt, writer.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (writerRepository *WriterRepository) Find(id string) (entity.Writer, error) {
	var writer entity.Writer

	rows, err := writerRepository.Db.Query("SELECT * FROM writers WHERE id = $1", id)
	if err != nil {
		return writer, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&writer.ID,
			&writer.Name,
			&writer.Picture,
			&writer.IsDeleted,
			&writer.CreatedAt,
			&writer.UpdatedAt,
			&writer.DeletedAt); err != nil {
			return writer, err
		}
	}

	if err = rows.Err(); err != nil {
		return writer, err
	}

	return writer, nil
}

func (writerRepository *WriterRepository) Update(writer *entity.Writer) error {
	stmt, err := writerRepository.Db.Prepare("UPDATE writers SET name = $1, picture = $2, is_deleted = $3, created_at = $4, updated_at = $5, deleted_at = $6 WHERE id = $7")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(writer.Name, writer.Picture, writer.IsDeleted, writer.CreatedAt, writer.UpdatedAt, writer.DeletedAt, writer.ID)
	if err != nil {
		return err
	}

	return nil
}

func (writerRepository *WriterRepository) Delete(writer *entity.Writer) error {
	stmt, err := writerRepository.Db.Prepare("UPDATE writers SET is_deleted = $1, deleted_at = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(writer.IsDeleted, writer.DeletedAt, writer.ID)
	if err != nil {
		return err
	}

	return nil
}

func (writerRepository *WriterRepository) IsDeleted(id string) error {
	var writer entity.Writer

	rows, err := writerRepository.Db.Query("SELECT * FROM writers WHERE id = $1", id)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.Scan(&writer.ID, &writer.Name, &writer.Picture, &writer.IsDeleted, &writer.CreatedAt, &writer.UpdatedAt, &writer.DeletedAt); err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (writerRepository *WriterRepository) FindAll() ([]entity.Writer, error) {
	rows, err := writerRepository.Db.Query("SELECT writer_id, name, picture, is_deleted, created_at, updated_at, deleted_at FROM writers")
	if err != nil {
		return nil, err
	}

	var writers []entity.Writer

	for rows.Next() {
		var writer entity.Writer

		if err := rows.Scan(&writer.ID, &writer.Name, &writer.Picture, &writer.IsDeleted, &writer.CreatedAt, &writer.UpdatedAt, &writer.DeletedAt); err != nil {
			return writers, err
		}

		writers = append(writers, writer)
	}

	if err = rows.Err(); err != nil {
		return writers, err
	}

	return writers, nil
}
