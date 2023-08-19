package v1

import (
	"net/http"

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
