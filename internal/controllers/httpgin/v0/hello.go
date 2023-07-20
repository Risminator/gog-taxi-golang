package v0

import (
	"net/http"

	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

type helloRoutes struct {
	hu usecase.Hello
}

func registerHelloRoutes(handler *gin.RouterGroup, hu usecase.Hello) {
	r := &helloRoutes{hu}

	h := handler.Group("/hello")
	{
		h.GET("/:name", r.sayHello)
	}
}

func (r *helloRoutes) sayHello(c *gin.Context) {
	msg, err := r.hu.SayHello(c.Param("name"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, sayHelloResponse{msg})
}
