package usecase

import (
	"context"

	"github.com/rafliputraa/petstore/internal/entity"
)

type (
	Customer interface {
		GetCustomer(context.Context) ([]entity.CustomerResponseDTO, error)
		GetCustomerById(c context.Context, id uint64) (*entity.CustomerResponseDTO, error)
		PostCustomer(context.Context, entity.CustomerRequestDTO) error
	}

	// CustomerRepo -.
	CustomerRepo interface {
		GetCustomer(context.Context) ([]entity.CustomerResponseDTO, error)
		GetCustomerById(c context.Context, id uint64) (*entity.CustomerResponseDTO, error)
		InsertCustomer(context.Context, entity.CustomerRequestDTO) error
	}
)
