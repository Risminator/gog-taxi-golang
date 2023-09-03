package v1

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TaxiRequestWsGateway interface {
	SendNewTaxiRequest(req model.TaxiRequest) error
	ConnectWebsocket(w http.ResponseWriter, r *http.Request, userId int, role model.UserRole, ct model.WebsocketClientType, reqId int, initEvent *model.Event) error
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
		h.GET("/stream-order-status/:requestId", routes.streamOrderStatus)
		h.GET("/stream-order-offer/:driverId", routes.streamOrderOffer)
		h.GET("/user/:id", routes.getRequestByUserId)
	}
}

func (r *taxiRequestRoutes) getRequestByUserId(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	role, err := model.UserRoleFromString(c.Query("role"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	msg, err := r.taxiUsecase.GetRequestByUserId(userId, role)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	if msg.TaxiRequestId == 0 {
		c.AbortWithError(http.StatusNotFound, errors.New("No active taxi request for user"))
	} else {
		c.JSON(http.StatusOK, *msg)
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
	c.JSON(http.StatusOK, *msg)
}

func (r *taxiRequestRoutes) getRequestsByStatus(c *gin.Context) {
	status, err := model.TaxiRequestStatusFromString(c.Param("status"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
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

	req, err := r.taxiUsecase.CreateRequest(body.TaxiRequestId, body.CustomerId, body.DriverId, body.DepartureId, body.DestinationId, body.Price)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	r.taxiWebsockets.SendNewTaxiRequest(*req)
	c.JSON(http.StatusCreated, req)
}

func (r *taxiRequestRoutes) streamOrderStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("requestId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	role, err := model.UserRoleFromString(c.Query("role"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	req, err := r.taxiUsecase.GetRequestById(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var userId int
	var wsClientType model.WebsocketClientType
	if role == model.DriverRole {
		userId = req.DriverId
		wsClientType = model.DriverCurrentTaxiRequestInfo
	} else {
		userId = req.CustomerId
		wsClientType = model.CustomerCurrentTaxiRequestInfo
	}

	r.taxiWebsockets.ConnectWebsocket(c.Writer, c.Request, userId, role, wsClientType, req.TaxiRequestId, nil)
}

// ws://localhost/taxi-request/websocket/:requestId?role=driver
func (r *taxiRequestRoutes) streamOrderOffer(c *gin.Context) {
	driverId, err := strconv.Atoi(c.Param("driverId"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// get current untaken taxi requests
	untakenRequests, err := r.taxiUsecase.GetRequestsByStatus(model.FindingDriver)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// prepare initial message with current taxi requests
	var initMessage model.Event
	initMessage.Type = model.EventNewTaxiRequest
	data, err := json.Marshal(untakenRequests)
	if err != nil {
		log.Printf("failed to marshal request info: %v", err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	initMessage.Payload = data

	r.taxiWebsockets.ConnectWebsocket(c.Writer, c.Request, driverId, model.DriverRole, model.DriverGetOffers, 0, &initMessage)
}
