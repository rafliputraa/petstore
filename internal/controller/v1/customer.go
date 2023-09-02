package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entity "github.com/rafliputraa/petstore/internal/entity/customer"
	usecase "github.com/rafliputraa/petstore/internal/usecase/customer"
	"github.com/rafliputraa/petstore/pkg/logger"
)

type customerRoutes struct {
	t usecase.Customer
	l logger.Interface
}

func newCustomerRoutes(handler *gin.RouterGroup, t usecase.Customer, l logger.Interface) {
	r := &customerRoutes{t, l}

	h := handler.Group("/customer")
	{
		h.GET("", r.getCustomer)
		h.GET("/:id", r.getCustomerById)
		h.POST("", r.postCustomer)
		h.PUT("/:id", r.updateCustomer)
		h.DELETE("/:id", r.deleteCustomer)
	}
}

type getCustomerResponse struct {
	Customer []entity.CustomerResponseDTO `json:"data"`
}

func (r *customerRoutes) getCustomer(c *gin.Context) {
	customers, err := r.t.GetCustomer(c.Request.Context())
	if err != nil {
		r.l.Error(err, "http - v1 - getCustomer")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, getCustomerResponse{customers})
}

type getCustomerByIdResponse struct {
	Customer []entity.CustomerResponseDTO `json:"data"`
}

func (r *customerRoutes) getCustomerById(c *gin.Context) {
	customer, err := r.t.GetCustomerById(c.Request.Context(), c.Param("id"))
	if err != nil {
		r.l.Error(err, "http - v1 - getCustomerById")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, getCustomerByIdResponse{customer})
}

func (r *customerRoutes) postCustomer(c *gin.Context) {
	var request entity.CustomerRequestDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - postCustomer")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := r.t.InsertCustomer(
		c.Request.Context(),
		entity.CustomerRequestDTO{
			FirstName:   request.FirstName,
			LastName:    request.LastName,
			Email:       request.Email,
			PhoneNumber: request.PhoneNumber,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - postCustomer")
		errorResponse(c, http.StatusInternalServerError, "post customer service problems")
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (r *customerRoutes) updateCustomer(c *gin.Context) {
	var request entity.CustomerRequestDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - updateCustomer")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := r.t.UpdateCustomer(
		c.Request.Context(),
		c.Param("id"),
		entity.CustomerRequestDTO{
			FirstName:   request.FirstName,
			LastName:    request.LastName,
			Email:       request.Email,
			PhoneNumber: request.PhoneNumber,
		},
	)
	if err != nil {
		r.l.Error(err, "http - v1 - updateCustomer")
		errorResponse(c, http.StatusInternalServerError, "update customer service problems")
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (r *customerRoutes) deleteCustomer(c *gin.Context) {
	err := r.t.DeleteCustomer(c.Request.Context(), c.Param("id"))
	if err != nil {
		r.l.Error(err, "http - v1 - deleteCustomer")
		errorResponse(c, http.StatusInternalServerError, "delete customer database problems")

		return
	}

	c.JSON(http.StatusNoContent, nil)
}
