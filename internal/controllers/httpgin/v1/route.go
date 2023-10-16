package v1

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/gin-gonic/gin"
)

type routeInfoRoutes struct {
	routeUsecase usecase.Route
}

func registerRouteInfoRoutes(r *gin.RouterGroup, ru usecase.Route) {
	routes := &routeInfoRoutes{ru}

	h := r.Group("route")
	{
		h.GET("/", routes.getRouteInfo)
	}
}

func parseLocationStr(str string) *model.Location {
	arr := strings.Split(str, ",")
	lon, err := strconv.ParseFloat(arr[0], 64)
	if err != nil {
		return nil
	}
	lat, err := strconv.ParseFloat(arr[1], 64)
	if err != nil {
		return nil
	}
	return &model.Location{
		Latitude:  lat,
		Longitude: lon,
	}
}

func (r *routeInfoRoutes) getRouteInfo(c *gin.Context) {
	profile := c.Query("profile")
	if profile == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("Profile parameter required"))
		return
	}

	lonlatsArr := strings.Split(c.Query("lonlats"), "|")

	start := parseLocationStr(lonlatsArr[0])
	end := parseLocationStr(lonlatsArr[1])

	feature, err := r.routeUsecase.GetRouteInfo(*start, *end, profile)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, feature)
}
