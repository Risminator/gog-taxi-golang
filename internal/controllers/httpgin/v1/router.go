package v1

import (
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.RouterGroup, cu usecase.Customer, du usecase.Dock, dru usecase.Driver, ru usecase.TaxiRequest, rws TaxiRequestWsGateway) {
	h := handler.Group("v1")
	{
		registerCustomerRoutes(h, cu)
		registerDockRoutes(h, du)
		registerDriverRoutes(h, dru)
		registerTaxiRequestRoutes(h, ru, rws)
	}
}
