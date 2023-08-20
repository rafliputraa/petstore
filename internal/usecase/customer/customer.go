package usecase

import (
	"context"
	"fmt"

	entity "github.com/rafliputraa/petstore/internal/entity"
)

type CustomerUseCase struct {
	repo CustomerRepo
}

// New -.
func New(r CustomerRepo) *CustomerUseCase {
	return &CustomerUseCase{
		repo: r,
	}
}

// Get Customer - retrieve all customer data.
func (uc *CustomerUseCase) GetCustomer(ctx context.Context) ([]entity.CustomerResponseDTO, error) {
	customers, err := uc.repo.GetCustomer(ctx)
	if err != nil {
		return nil, fmt.Errorf("CustomerUseCase - GetCustomer - s.repo.GetCustomer: %w", err)
	}

	return customers, nil
}

// Get Customer - retrieve all customer data.
func (uc *CustomerUseCase) GetCustomerById(ctx context.Context, id uint64) (*entity.CustomerResponseDTO, error) {
	customer, err := uc.repo.GetCustomerById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("CustomerUseCase - GetCustomerById - s.repo.GetCustomerById: %w", err)
	}

	return customer, nil
}

// Post Customer - insert a customer data.
func (uc *CustomerUseCase) PostCustomer(ctx context.Context, t entity.CustomerRequestDTO) error {
	err := uc.repo.InsertCustomer(ctx, t)
	if err != nil {
		return fmt.Errorf("CustomerUseCase - PostCustomer - s.repo.PostCustomer: %w", err)
	}

	return nil
}
