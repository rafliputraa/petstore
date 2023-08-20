package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rafliputraa/petstore/internal/entity"
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
	}
}

type getCustomerResponse struct {
	Customer []entity.CustomerResponseDTO `json:"customer"`
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

func (r *customerRoutes) getCustomerById(c *gin.Context) {
	uint64Id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	customer, err := r.t.GetCustomerById(c.Request.Context(), uint64Id)
	if err != nil {
		r.l.Error(err, "http - v1 - getCustomerById")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	c.JSON(http.StatusOK, &customer)
}

func (r *customerRoutes) postCustomer(c *gin.Context) {
	var request entity.CustomerRequestDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		r.l.Error(err, "http - v1 - postCustomer")
		errorResponse(c, http.StatusBadRequest, "invalid request body")
		return
	}

	err := r.t.PostCustomer(
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
