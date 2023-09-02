package usecase

import (
	"context"

	entity "github.com/rafliputraa/petstore/internal/entity/customer"
)

type (
	Customer interface {
		GetCustomer(context.Context) ([]entity.CustomerResponseDTO, error)
		GetCustomerById(c context.Context, id string) ([]entity.CustomerResponseDTO, error)
		InsertCustomer(context.Context, entity.CustomerRequestDTO) error
		UpdateCustomer(c context.Context, id string, e entity.CustomerRequestDTO) error
		DeleteCustomer(c context.Context, id string) error
	}

	// CustomerRepo -.
	CustomerRepo interface {
		GetCustomer(context.Context) ([]entity.CustomerResponseDTO, error)
		GetCustomerById(c context.Context, id string) ([]entity.CustomerResponseDTO, error)
		InsertCustomer(context.Context, entity.CustomerRequestDTO) error
		UpdateCustomer(c context.Context, id string, e entity.CustomerRequestDTO) error
		DeleteCustomer(c context.Context, id string) error
	}
)
