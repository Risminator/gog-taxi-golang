package v0

import (
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.RouterGroup, a usecase.Hello) {
	handler.Use(gin.Recovery())
	handler.Use(gin.Logger())

	h := handler.Group("/v0")
	{
		registerHelloRoutes(h, a)
	}
}
