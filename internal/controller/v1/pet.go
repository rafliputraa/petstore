package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entity "github.com/rafliputraa/petstore/internal/entity/pet"
	usecase "github.com/rafliputraa/petstore/internal/usecase/pet"
	"github.com/rafliputraa/petstore/pkg/logger"
)

type petRoutes struct {
	t usecase.Pet
	l logger.Interface
}

func newPetRoutes(handler *gin.RouterGroup, t usecase.Pet, l logger.Interface) {
	r := &petRoutes{t, l}

	h := handler.Group("/pet")
	{
		h.GET("", r.getPet)
		h.GET("/:id", r.getPetById)
		h.POST("", r.postPet)
		h.PUT("/:id", r.updatePet)
		h.DELETE("/:id", r.deletePet)
	}
}

type getPetResponse struct {
	Pet []entity.PetResponseDTO `json:"data"`
}

func (r *petRoutes) getPet(c *gin.Context) {
	pets, err := r.t.GetPet(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - getPet")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, getPetResponse{pets})
}

type getPetByIdResponse struct {
	Pet []entity.PetResponseDTO `json:"data"`
}

func (r *petRoutes) getPetById(c *gin.Context) {
	pet, err := r.t.GetPetById(c.Request.Context(), c.Param("id"))
	if err != nil {
		r.l.Error(err, "http - v1 - getPetById")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, getPetByIdResponse{pet})
}

func (r *petRoutes) postPet(c *gin.Context) {
	var request entity.PetRequestDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - postPet")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := r.t.InsertPet(
		c.Request.Context(),
		entity.PetRequestDTO{
			Name:      request.Name,
			Species:   request.Species,
			Age:       request.Age,
			Available: request.Available,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - postPet")
		errorResponse(c, http.StatusInternalServerError, "post pet service problems")
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (r *petRoutes) updatePet(c *gin.Context) {
	var request entity.PetRequestDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - updatePet")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := r.t.UpdatePet(
		c.Request.Context(),
		c.Param("id"),
		entity.PetRequestDTO{
			Name:      request.Name,
			Species:   request.Species,
			Age:       request.Age,
			Available: request.Available,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - updatePet")
		errorResponse(c, http.StatusInternalServerError, "update pet service problems")
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (r *petRoutes) deletePet(c *gin.Context) {
	err := r.t.DeletePet(c.Request.Context(), c.Param("id"))
	if err != nil {
		r.l.Error(err, "http - v1 - deletePet")
		errorResponse(c, http.StatusInternalServerError, "delete pet database problems")

		return
	}

	c.JSON(http.StatusNoContent, nil)
}
