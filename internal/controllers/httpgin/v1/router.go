package v1

import (
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.RouterGroup, cu usecase.Customer) {
	h := handler.Group("v1")
	{
		registerCustomerRoutes(h, cu)
	}
}
