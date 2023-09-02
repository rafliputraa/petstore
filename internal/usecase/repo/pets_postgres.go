package repo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	entity "github.com/rafliputraa/petstore/internal/entity/pet"
	"github.com/rafliputraa/petstore/pkg/postgres"
)

// PetRepo -.
type PetRepo struct {
	*postgres.Postgres
}

// New -.
func NewPetRepo(pg *postgres.Postgres) *PetRepo {
	return &PetRepo{pg}
}

// GetPet -.
func (r *PetRepo) GetPet(ctx context.Context) ([]entity.PetResponseDTO, error) {
	sql, _, err := r.Builder.
		Select("pet_id, name, species, age, available").
		From("pets").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PetRepo - GetPet - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("PetRepo - GetPet - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.PetResponseDTO, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.PetResponseDTO{}

		err = rows.Scan(&e.PetId, &e.Name, &e.Species, &e.Age, &e.Available)
		if err != nil {
			return nil, fmt.Errorf("PetRepo - GetPet - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// GetPetById -.
func (r *PetRepo) GetPetById(ctx context.Context, id string) ([]entity.PetResponseDTO, error) {
	sql, _, err := r.Builder.
		Select("pet_id, name, species, age, available").
		From("pets").
		Where("pet_id = $1").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("PetRepo - GetPetById - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, id)
	if err != nil {
		return nil, fmt.Errorf("PetRepo - GetPetById - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.PetResponseDTO, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.PetResponseDTO{}

		err = rows.Scan(&e.PetId, &e.Name, &e.Species, &e.Age, &e.Available)
		if err != nil {
			return nil, fmt.Errorf("PetRepo - GetPetById - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// InsertPet -.
func (r *PetRepo) InsertPet(ctx context.Context, t entity.PetRequestDTO) error {
	sql, args, err := r.Builder.
		Insert("pets").
		Columns("name, species, age, available").
		Values(t.Name, t.Species, t.Age, t.Available).
		ToSql()
	if err != nil {
		return fmt.Errorf("PetRepo - Insert - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PetRepo - Insert - r.Pool.Exec: %w", err)
	}

	return nil
}

// UpdatePet -.
func (r *PetRepo) UpdatePet(ctx context.Context, id string, t entity.PetRequestDTO) error {
	sql, args, err := r.Builder.
		Update("pets").
		Set("name", t.Name).
		Set("species", t.Species).
		Set("age", t.Age).
		Set("available", t.Available).
		Where(squirrel.Eq{"pet_id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("PetRepo - Update - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("PetRepo - Update - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeletePet -.
func (r *PetRepo) DeletePet(ctx context.Context, id string) error {
	sql, _, err := r.Builder.
		Delete("pets").
		Where("pet_id = $1").
		ToSql()
	if err != nil {
		return fmt.Errorf("PetRepo - Delete - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("PetRepo - Delete - r.Pool.Exec: %w", err)
	}

	return nil
}
