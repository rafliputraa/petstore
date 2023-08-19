package usecase

import (
	"context"

	"github.com/rafliputraa/petstore/internal/entity"
)

type (
	Customer interface {
		GetCustomer(context.Context) ([]entity.CustomerResponseDTO, error)
		PostCustomer(context.Context)
	}

	// CustomerRepo -.
	CustomerRepo interface {
		GetCustomer(context.Context) ([]entity.CustomerResponseDTO, error)
	}
)
