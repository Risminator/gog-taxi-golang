package v0

import (
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.RouterGroup, hu usecase.Hello) {
	h := handler.Group("/v0")
	{
		registerHelloRoutes(h, hu)
	}
}
