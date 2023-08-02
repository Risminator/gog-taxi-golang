package v1

import (
	"net/http"
	"strconv"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TaxiRequestWsGateway interface {
	SendNewTaxiRequest(req model.TaxiRequest) error
	ConnectWebsocket(w http.ResponseWriter, r *http.Request, u *model.User)
}

type taxiRequestRoutes struct {
	taxiUsecase    usecase.TaxiRequest
	taxiWebsockets TaxiRequestWsGateway
}

func registerTaxiRequestRoutes(r *gin.RouterGroup, taxiUsecase usecase.TaxiRequest, taxiWebsockets TaxiRequestWsGateway) {
	routes := &taxiRequestRoutes{taxiUsecase, taxiWebsockets}

	h := r.Group("/taxi-request")
	{
		h.GET("/:id", routes.getRequestById)
		h.GET("/status/:status", routes.getRequestsByStatus)
		h.POST("/", routes.createRequest)
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
	status, err := model.TaxiRequestStatusFromString(c.Param("status"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

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
	// TODO: add websocket functions
	c.JSON(http.StatusOK, msg)
}
