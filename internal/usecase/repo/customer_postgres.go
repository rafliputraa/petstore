package repo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/rafliputraa/petstore/internal/entity"
	"github.com/rafliputraa/petstore/pkg/postgres"
)

const _defaultEntityCap = 64

// CustomerRepo -.
type CustomerRepo struct {
	*postgres.Postgres
}

// New -.
func New(pg *postgres.Postgres) *CustomerRepo {
	return &CustomerRepo{pg}
}

// GetCustomer -.
func (r *CustomerRepo) GetCustomer(ctx context.Context) ([]entity.CustomerResponseDTO, error) {
	sql, _, err := r.Builder.
		Select("*").
		From("customers").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("CustomerRepo - GetCustomer - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("CustomerRepo - GetCustomer - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.CustomerResponseDTO, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.CustomerResponseDTO{}

		err = rows.Scan(&e.CustomerId, &e.FirstName, &e.LastName, &e.Email, &e.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("CustomerRepo - GetCustomer - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// GetCustomerById -.
func (r *CustomerRepo) GetCustomerById(ctx context.Context, id string) ([]entity.CustomerResponseDTO, error) {
	sql, _, err := r.Builder.
		Select("*").
		From("customers").
		Where("customer_id = $1").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("CustomerRepo - GetCustomerById - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, id)
	if err != nil {
		return nil, fmt.Errorf("CustomerRepo - GetCustomerById - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.CustomerResponseDTO, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.CustomerResponseDTO{}

		err = rows.Scan(&e.CustomerId, &e.FirstName, &e.LastName, &e.Email, &e.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("CustomerRepo - GetCustomerById - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}

// InsertCustomer -.
func (r *CustomerRepo) InsertCustomer(ctx context.Context, t entity.CustomerRequestDTO) error {
	sql, args, err := r.Builder.
		Insert("customers").
		Columns("first_name, last_name, email, phone").
		Values(t.FirstName, t.LastName, t.Email, t.PhoneNumber).
		ToSql()
	if err != nil {
		return fmt.Errorf("CustomerRepo - Insert - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("CustomerRepo - Insert - r.Pool.Exec: %w", err)
	}

	return nil
}

// UpdateCustomer -.
func (r *CustomerRepo) UpdateCustomer(ctx context.Context, id string, t entity.CustomerRequestDTO) error {
	sql, args, err := r.Builder.
		Update("customers").
		Set("first_name", t.FirstName).
		Set("last_name", t.LastName).
		Set("email", t.Email).
		Set("phone", t.PhoneNumber).
		Where(squirrel.Eq{"customer_id": id}).
		ToSql()
	if err != nil {
		return fmt.Errorf("CustomerRepo - Update - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("CustomerRepo - Update - r.Pool.Exec: %w", err)
	}

	return nil
}

// DeleteCustomer -.
func (r *CustomerRepo) DeleteCustomer(ctx context.Context, id string) error {
	sql, _, err := r.Builder.
		Delete("customers").
		Where("customer_id = $1").
		ToSql()
	if err != nil {
		return fmt.Errorf("CustomerRepo - Delete - r.Builder: %w", err)
	}

	_, err = r.Pool.Exec(ctx, sql, id)
	if err != nil {
		return fmt.Errorf("CustomerRepo - Delete - r.Pool.Exec: %w", err)
	}

	return nil
}
