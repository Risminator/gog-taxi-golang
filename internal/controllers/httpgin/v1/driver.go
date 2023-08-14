package v1

import (
	"net/http"
	"strconv"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

type driverInput struct {
	FirstName    string             `json:"first_name" binding:"required"`
	LastName     string             `json:"last_name" binding:"required"`
	VesselId     int                `json:"vessel_id" binding:"required"`
	Status       model.DriverStatus `json:"status" binding:"required"`
	Balance      float64            `json:"balance" binding:"required"`
	CertFirstAid int                `json:"cert_first_aid" binding:"required"`
	CertDriving  int                `json:"cert_driving" binding:"required"`
}

func registerDriverRoutes(handler *gin.RouterGroup, du usecase.Driver) {
	r := &driverRoutes{du}

	// what to do with names?
	h := handler.Group("/driver")
	{
		h.GET("/", r.getAllDrivers)
		h.GET("/:id", r.getDriverByID)
		h.POST("/", r.createDriver)
		h.PUT("/:id", r.updateDriver)
		h.DELETE("/:id", r.deleteDriver)
	}
}

type driverRoutes struct {
	du usecase.Driver
}

func (r *driverRoutes) getAllDrivers(c *gin.Context) {
	msg, err := r.du.GetAllDrivers()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (r *driverRoutes) getDriverByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.du.GetDriverByID(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}

func (r *driverRoutes) createDriver(c *gin.Context) {
	var body driverInput
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.du.CreateDriver(body.FirstName, body.LastName, body.VesselId, body.CertFirstAid, body.CertDriving)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}

func (r *driverRoutes) updateDriver(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var body driverInput
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.du.UpdateDriver(id, body.FirstName, body.LastName, body.VesselId, body.Status, body.Balance, body.CertFirstAid, body.CertDriving)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}

func (r *driverRoutes) deleteDriver(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.du.DeleteDriver(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}
