package v1

import (
	"net/http"
	"strconv"

	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

type customerRoutes struct {
	cu usecase.Customer
}

func registerCustomerRoutes(handler *gin.RouterGroup, cu usecase.Customer) {
	r := &customerRoutes{cu}

	h := handler.Group("/")
	{
		h.GET("/customer", r.getAllCustomers)
		h.GET("/customer/:id", r.getCustomerByID)
	}
}

func (r *customerRoutes) getAllCustomers(c *gin.Context) {
	msg, err := r.cu.GetAllCustomers()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (r *customerRoutes) getCustomerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	msg, err := r.cu.GetCustomerByID(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, *msg)
}
