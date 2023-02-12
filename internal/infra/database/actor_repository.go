package database

import (
	"database/sql"

	"github.com/GuilhermeDeOliveiraAmorim/you-choose/internal/entity"
)

type ActorRepository struct {
	Db *sql.DB
}

func NewActorRepository(db *sql.DB) *ActorRepository {
	return &ActorRepository{
		Db: db,
	}
}

func (actorRepository *ActorRepository) Create(actor *entity.Actor) error {
	stmt, err := actorRepository.Db.Prepare("INSERT INTO actors (actor_id, name, picture, is_deleted, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(actor.ID, actor.Name, actor.Picture, actor.IsDeleted, actor.CreatedAt, actor.UpdatedAt, actor.DeletedAt)
	if err != nil {
		return err
	}

	return nil
}

func (actorRepository *ActorRepository) Find(id string) (entity.Actor, error) {
	var actor entity.Actor

	rows, err := actorRepository.Db.Query("SELECT * FROM actors WHERE id = $1", id)
	if err != nil {
		return actor, err
	}

	for rows.Next() {
		if err := rows.Scan(
			&actor.ID,
			&actor.Name,
			&actor.Picture,
			&actor.IsDeleted,
			&actor.CreatedAt,
			&actor.UpdatedAt,
			&actor.DeletedAt); err != nil {
			return actor, err
		}
	}

	if err = rows.Err(); err != nil {
		return actor, err
	}

	return actor, nil
}

func (actorRepository *ActorRepository) Update(actor *entity.Actor) error {
	stmt, err := actorRepository.Db.Prepare("UPDATE actors SET name = $1, picture = $2, is_deleted = $3, created_at = $4, updated_at = $5, deleted_at = $6 WHERE id = $7")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(actor.Name, actor.Picture, actor.IsDeleted, actor.CreatedAt, actor.UpdatedAt, actor.DeletedAt, actor.ID)
	if err != nil {
		return err
	}

	return nil
}

func (actorRepository *ActorRepository) Delete(actor *entity.Actor) error {
	stmt, err := actorRepository.Db.Prepare("UPDATE actors SET is_deleted = $1, deleted_at = $2 WHERE id = $3")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(actor.IsDeleted, actor.DeletedAt, actor.ID)
	if err != nil {
		return err
	}

	return nil
}

func (actorRepository *ActorRepository) IsDeleted(id string) error {
	var actor entity.Actor

	rows, err := actorRepository.Db.Query("SELECT * FROM actors WHERE id = $1", id)
	if err != nil {
		return err
	}

	for rows.Next() {
		if err := rows.Scan(&actor.ID, &actor.Name, &actor.Picture, &actor.IsDeleted, &actor.CreatedAt, &actor.UpdatedAt, &actor.DeletedAt); err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

func (actorRepository *ActorRepository) FindAll() ([]entity.Actor, error) {
	rows, err := actorRepository.Db.Query("SELECT actor_id, name, picture, is_deleted, created_at, updated_at, deleted_at FROM actors")
	if err != nil {
		return nil, err
	}

	var actors []entity.Actor

	for rows.Next() {
		var actor entity.Actor

		if err := rows.Scan(&actor.ID, &actor.Name, &actor.Picture, &actor.IsDeleted, &actor.CreatedAt, &actor.UpdatedAt, &actor.DeletedAt); err != nil {
			return actors, err
		}

		actors = append(actors, actor)
	}

	if err = rows.Err(); err != nil {
		return actors, err
	}

	return actors, nil
}
