package repo

import (
	"context"
	"fmt"
	"log"

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
func (r *CustomerRepo) GetCustomerById(ctx context.Context, id uint64) (*entity.CustomerResponseDTO, error) {
	// strValue := strconv.FormatUint(id, 10)
	e := entity.CustomerResponseDTO{}
	sql, _, err := r.Builder.
		Select("*").
		From("customers").
		Where("customer_id = $1").
		ToSql()
	log.Println(sql)
	if err != nil {
		return nil, fmt.Errorf("CustomerRepo - GetCustomerById - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql, id)
	if err != nil {
		return nil, fmt.Errorf("CustomerRepo - GetCustomerById - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	err = rows.Scan(
		&e.CustomerId,
		&e.FirstName,
		&e.LastName,
		&e.Email,
		&e.PhoneNumber,
	)
	if err != nil {
		return nil, fmt.Errorf("CustomerRepo - GetCustomerById - rows.Scan: %w", err)
	}

	return &e, nil
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
