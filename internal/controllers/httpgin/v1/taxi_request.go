package v1

import (
	"net/http"
	"strconv"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

type taxiRequestRoutes struct {
	taxiUsecase usecase.TaxiRequest
}

func registerTaxiRequestRoutes(r *gin.RouterGroup, taxiUsecase usecase.TaxiRequest) {
	routes := &taxiRequestRoutes{taxiUsecase}

	h := r.Group("/taxi-request")
	{
		h.GET("/:id", routes.getRequestById)
		h.GET("status/:status", routes.getRequestsByStatus)
		h.POST("/add", routes.createRequest)
	}
}

func (r *taxiRequestRoutes) getRequestById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.taxiUsecase.GetRequestById(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (r *taxiRequestRoutes) getRequestsByStatus(c *gin.Context) {
	status := model.ParseTaxiRequestStatus(c.Param("status"))

	msg, err := r.taxiUsecase.GetRequestsByStatus(status)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
}

func (r *taxiRequestRoutes) createRequest(c *gin.Context) {
	var body model.TaxiRequest
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.taxiUsecase.CreateRequest(body.TaxiRequestId, body.ClientId, body.DriverId, body.DepartureId, body.DestinationId, body.Price)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, msg)
}