package usecase

import (
	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/paulmach/orb/geojson"
)

type Route interface {
	GetRouteInfo(start model.Location, end model.Location, profile string) (*geojson.FeatureCollection, error)
}
type RouteWebApi interface {
	GetRouteInfo(start model.Location, end model.Location, profile string) (*geojson.FeatureCollection, error)
}

type routeUsecase struct {
	webApi RouteWebApi
}

func NewRouteUsecase(webApi RouteWebApi) Route {
	return &routeUsecase{webApi}
}

// GetRouteInfo implements Route.
func (ru *routeUsecase) GetRouteInfo(start model.Location, end model.Location, profile string) (*geojson.FeatureCollection, error) {
	feature, err := ru.webApi.GetRouteInfo(start, end, profile)
	if err != nil {
		return nil, err
	}

	return feature, nil
}
