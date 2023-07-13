package httpgin

import (
	"net/http"

	"github.com/Risminator/gog-taxi-golang/internal/app"
	"github.com/gin-gonic/gin"
)

func SayHello(a app.App) gin.HandlerFunc {
	fn := func(c *gin.Context) {
		msg, err := a.SayHello(c.Param("name"))
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		c.JSON(http.StatusOK, sayHelloResponse{msg})
	}

	return gin.HandlerFunc(fn)
}
