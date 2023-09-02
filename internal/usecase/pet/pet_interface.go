package usecase

import (
	"context"

	entity "github.com/rafliputraa/petstore/internal/entity/pet"
)

type (
	Pet interface {
		GetPet(context.Context) ([]entity.PetResponseDTO, error)
		GetPetById(c context.Context, id string) ([]entity.PetResponseDTO, error)
		InsertPet(context.Context, entity.PetRequestDTO) error
		UpdatePet(c context.Context, id string, e entity.PetRequestDTO) error
		DeletePet(c context.Context, id string) error
	}

	// PetRepo -.
	PetRepo interface {
		GetPet(context.Context) ([]entity.PetResponseDTO, error)
		GetPetById(c context.Context, id string) ([]entity.PetResponseDTO, error)
		InsertPet(context.Context, entity.PetRequestDTO) error
		UpdatePet(c context.Context, id string, e entity.PetRequestDTO) error
		DeletePet(c context.Context, id string) error
	}
)
