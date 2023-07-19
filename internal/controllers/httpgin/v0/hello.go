package v0

import (
	"net/http"

	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

type helloRoutes struct {
	a usecase.Hello
}

func registerHelloRoutes(handler *gin.RouterGroup, a usecase.Hello) {
	r := &helloRoutes{a}

	h := handler.Group("/")
	{
		h.GET("/:name", r.sayHello)
	}
}

func (r *helloRoutes) sayHello(c *gin.Context) {
	msg, err := r.a.SayHello(c.Param("name"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, sayHelloResponse{msg})
}
