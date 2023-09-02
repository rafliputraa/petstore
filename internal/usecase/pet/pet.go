package usecase

import (
	"context"
	"fmt"

	entity "github.com/rafliputraa/petstore/internal/entity/pet"
)

type PetUseCase struct {
	repo PetRepo
}

// New -.
func New(r PetRepo) *PetUseCase {
	return &PetUseCase{
		repo: r,
	}
}

// Get Pet - retrieve all pet data.
func (uc *PetUseCase) GetPet(ctx context.Context) ([]entity.PetResponseDTO, error) {
	pets, err := uc.repo.GetPet(ctx)
	if err != nil {
		return nil, fmt.Errorf("PetUseCase - GetPet - s.repo.GetPet: %w", err)
	}

	return pets, nil
}

// Get Pet - retrieve all pet data.
func (uc *PetUseCase) GetPetById(ctx context.Context, id string) ([]entity.PetResponseDTO, error) {
	pet, err := uc.repo.GetPetById(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("PetUseCase - GetPetById - s.repo.GetPetById: %w", err)
	}

	return pet, nil
}

// Post Pet - insert a pet data.
func (uc *PetUseCase) InsertPet(ctx context.Context, t entity.PetRequestDTO) error {
	err := uc.repo.InsertPet(ctx, t)
	if err != nil {
		return fmt.Errorf("PetUseCase -InsertPet - s.repo.PostPet: %w", err)
	}

	return nil
}

// Update Pet - update a pet data.
func (uc *PetUseCase) UpdatePet(ctx context.Context, id string, t entity.PetRequestDTO) error {
	err := uc.repo.UpdatePet(ctx, id, t)
	if err != nil {
		return fmt.Errorf("PetUseCase - UpdatePet - s.repo.UpdatePet: %w", err)
	}

	return nil
}

// Delete Pet - delete a pet data.
func (uc *PetUseCase) DeletePet(ctx context.Context, id string) error {
	err := uc.repo.DeletePet(ctx, id)
	if err != nil {
		return fmt.Errorf("PetUseCase - DeletePet - s.repo.DeletePet: %w", err)
	}

	return nil
}
