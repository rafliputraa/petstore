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

// Get Customer - getting all customer data from customer.
func (uc *CustomerUseCase) GetCustomer(ctx context.Context) ([]entity.CustomerResponseDTO, error) {
	customers, err := uc.repo.GetCustomer(ctx)
	if err != nil {
		return nil, fmt.Errorf("CustomerUseCase - GetCustomer - s.repo.GetCustomer: %w", err)
	}

	return customers, nil
}
