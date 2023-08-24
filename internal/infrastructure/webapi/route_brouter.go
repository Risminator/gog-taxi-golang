package webapi

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Risminator/gog-taxi-golang/internal/domain/model"
	"github.com/Risminator/gog-taxi-golang/internal/usecase"
	"github.com/paulmach/orb/geojson"
)

type routerBrouter struct {
}

func NewRouterBrouter() usecase.RouteWebApi {
	return &routerBrouter{}
}

// GetRouteInfo implements usecase.RouteWebApi.
func (*routerBrouter) GetRouteInfo(start model.Location, end model.Location) (*geojson.FeatureCollection, error) {
	req, err := http.NewRequest("GET", "http://localhost:17777/brouter", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("lonlats", fmt.Sprintf("%g,%g|%g,%g", start.Longitude, start.Latitude, end.Longitude, end.Latitude))
	q.Add("profile", "river")
	q.Add("format", "geojson")
	q.Add("alternativeidx", "0")
	queryURL := strings.Replace(q.Encode(), "%2C", ",", -1)
	queryURL = strings.Replace(queryURL, "%7C", "|", -1)
	req.URL.RawQuery = queryURL

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	feature, err := geojson.UnmarshalFeatureCollection(body)
	if err != nil {
		return nil, err
	}

	return feature, nil
}
