package v1

import (
	"net/http"
	"strconv"

	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

type customerFullInput struct {
	Phone     string `json:"phone" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

type customerPartialInput struct {
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type customerRoutes struct {
	cu usecase.Customer
}

func registerCustomerRoutes(handler *gin.RouterGroup, cu usecase.Customer) {
	r := &customerRoutes{cu}

	// what to do with names?
	h := handler.Group("/customer")
	{
		h.GET("/", r.getAllCustomers)
		h.GET("/:id", r.getCustomerByID)
		h.POST("/", r.createCustomer)
		h.PUT("/:id", r.updateCustomer)
		h.PATCH("/:id", r.updateCustomerPartially)
		h.DELETE("/:id", r.deleteCustomer)
	}
}

func (r *customerRoutes) getAllCustomers(c *gin.Context) {
	msg, err := r.cu.GetAllCustomers()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (r *customerRoutes) getCustomerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.cu.GetCustomerByID(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}

func (r *customerRoutes) createCustomer(c *gin.Context) {
	var body customerFullInput
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.cu.CreateCustomer(body.Phone, body.FirstName, body.LastName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}

func (r *customerRoutes) updateCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var body customerFullInput
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.cu.UpdateCustomer(id, body.Phone, body.FirstName, body.LastName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}

func (r *customerRoutes) updateCustomerPartially(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var body customerPartialInput
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.cu.UpdateCustomer(id, body.Phone, body.FirstName, body.LastName)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}

func (r *customerRoutes) deleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.cu.DeleteCustomer(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}
