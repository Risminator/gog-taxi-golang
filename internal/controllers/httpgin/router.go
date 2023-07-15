package httpgin

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/Risminator/gog-taxi-golang/internal/usecase/app"
	"github.com/gin-gonic/gin"
)

func AppRouter(r *gin.RouterGroup, a app.App) {
	//r.Use(gin.CustomRecovery(CustomPanicRecover))
	r.Use(gin.Recovery())
	//r.Use(CustomLogger)
	r.Use(gin.Logger())

	r.GET("/:name", SayHello(a))
}

func CustomLogger(c *gin.Context) {
	t := time.Now().UTC()
	c.Next()
	latency := time.Since(t)
	status := c.Writer.Status()

	log.Println("latency", latency, "method", c.Request.Method, "path", c.Request.URL.Path, "status", status)
}

func CustomPanicRecover(c *gin.Context, err any) {
	log.Println("panic: " + err.(error).Error())
	buf := make([]byte, 2048)
	n := runtime.Stack(buf, false)
	log.Println(string(buf[:n]))
	c.AbortWithStatusJSON(http.StatusInternalServerError, err.(error).Error())
}
