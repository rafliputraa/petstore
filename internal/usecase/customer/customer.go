package usecase

import (
	"context"
	"fmt"

	entity "github.com/rafliputraa/petstore/internal/entity/customer"
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
func (uc *CustomerUseCase) GetCustomerById(ctx context.Context, id string) ([]entity.CustomerResponseDTO, error) {
	customer, err := uc.repo.GetCustomerById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("CustomerUseCase - GetCustomerById - s.repo.GetCustomerById: %w", err)
	}

	return customer, nil
}

// Post Customer - insert a customer data.
func (uc *CustomerUseCase) InsertCustomer(ctx context.Context, t entity.CustomerRequestDTO) error {
	err := uc.repo.InsertCustomer(ctx, t)
	if err != nil {
		return fmt.Errorf("CustomerUseCase -InsertCustomer - s.repo.PostCustomer: %w", err)
	}

	return nil
}

// Update Customer - update a customer data.
func (uc *CustomerUseCase) UpdateCustomer(ctx context.Context, id string, t entity.CustomerRequestDTO) error {
	err := uc.repo.UpdateCustomer(ctx, id, t)
	if err != nil {
		return fmt.Errorf("CustomerUseCase - UpdateCustomer - s.repo.UpdateCustomer: %w", err)
	}

	return nil
}

// Delete Customer - delete a customer data.
func (uc *CustomerUseCase) DeleteCustomer(ctx context.Context, id string) error {
	err := uc.repo.DeleteCustomer(ctx, id)
	if err != nil {
		return fmt.Errorf("CustomerUseCase - DeleteCustomer - s.repo.DeleteCustomer: %w", err)
	}

	return nil
}
