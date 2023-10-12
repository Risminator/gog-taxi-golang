package v1

import (
	"net/http"
	"strconv"

	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

type vesselInput struct {
	Model      string  `json:"model"`
	Seats      int     `json:"seats"`
	IsApproved bool    `json:"isApproved"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

func registerVesselRoutes(handler *gin.RouterGroup, vu usecase.Vessel) {
	r := &vesselRoutes{vu}

	// what to do with names?
	h := handler.Group("/vessel")
	{
		h.GET("/", r.getAllVessels)
		h.GET("/:id", r.getVesselByID)
		h.GET("/request/:id", r.getVesselByRequestId)
		h.POST("/", r.createVessel)
		h.PUT("/:id", r.updateVessel)
		h.DELETE("/:id", r.deleteVessel)
	}
}

type vesselRoutes struct {
	vu usecase.Vessel
}

func (r *vesselRoutes) getAllVessels(c *gin.Context) {
	msg, err := r.vu.GetAllVessels()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (r *vesselRoutes) getVesselByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.vu.GetVesselByID(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}

func (r *vesselRoutes) getVesselByRequestId(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.vu.GetVesselByRequestId(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}

func (r *vesselRoutes) createVessel(c *gin.Context) {
	var body vesselInput
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.vu.CreateVessel(body.Model, body.Seats, body.IsApproved, body.Latitude, body.Longitude)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}

func (r *vesselRoutes) updateVessel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var body vesselInput
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.vu.UpdateVessel(id, body.Model, body.Seats, body.IsApproved, body.Latitude, body.Longitude)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}

func (r *vesselRoutes) deleteVessel(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.vu.DeleteVessel(id)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}
