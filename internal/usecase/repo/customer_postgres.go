package repo

import (
	"context"
	"fmt"

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
		Select("customer_id, first_name, last_name, email, phone").
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

		err = rows.Scan(&e.CustomerId, &e.Email, &e.FirstName, &e.LastName, &e.PhoneNumber)
		if err != nil {
			return nil, fmt.Errorf("CustomerRepo - GetCustomer - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}
