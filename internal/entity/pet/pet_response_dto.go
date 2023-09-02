// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

// Customer -.
type PetResponseDTO struct {
	PetId     uint64 `json:"pet_id"`
	Name      string `json:"name"`
	Species   string `json:"species"`
	Age       int    `json:"age"`
	Available bool   `json:"available"`
}
