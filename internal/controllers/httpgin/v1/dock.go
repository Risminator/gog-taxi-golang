package v1

import (
	"net/http"
	"strconv"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

type dockInput struct {
	Name      string  `json:"name" binding:"required"`
	Active    bool    `json:"active" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

func registerDockRoutes(r *gin.RouterGroup, du usecase.Dock) {
	routes := &dockRoutes{du}

	h := r.Group("/dock")
	{
		h.GET("/", routes.getDocks)
		h.GET("/:id", routes.getDockById)
		h.GET("/nearest", routes.getNearestDocks)
		h.POST("/", routes.createDock)
		h.PUT("/", routes.updateDock)
	}
}

type dockRoutes struct {
	du usecase.Dock
}

func (r *dockRoutes) getDocks(c *gin.Context) {
	msg, err := r.du.GetDocks()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (r *dockRoutes) getDockById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.du.GetDockById(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, *msg)
}

func (r *dockRoutes) getNearestDocks(c *gin.Context) {
	lat, err := strconv.ParseFloat(c.Query("latitude"), 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	lon, err := strconv.ParseFloat(c.Query("longitude"), 64)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	count := 1
	if query := c.Query("count"); query != "" {
		count, err = strconv.Atoi(query)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	msg, err := r.du.GetNearestDocks(lat, lon, count)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (r *dockRoutes) createDock(c *gin.Context) {
	var body dockInput
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.du.CreateDock(body.Name, body.Latitude, body.Longitude)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (r *dockRoutes) updateDock(c *gin.Context) {
	var body model.Dock
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.du.UpdateDock(body.DockId, body.Name, body.Latitude, body.Longitude)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
}
