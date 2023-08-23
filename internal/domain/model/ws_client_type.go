package model

import "errors"

type WebsocketClientType string

const (
	UnknownWsClientType            WebsocketClientType = ""
	CustomerCurrentTaxiRequestInfo WebsocketClientType = "customerCurrentTaxiRequestInfo"
	DriverCurrentTaxiRequestInfo   WebsocketClientType = "driverCurrentTaxiRequestInfo"
	DriverGetOffers                WebsocketClientType = "driverGetOffers"
)

func WsClientTypeFromString(str string) (WebsocketClientType, error) {
	switch str {
	case "customerCurrentTaxiRequestInfo":
		return CustomerCurrentTaxiRequestInfo, nil
	case "driverCurrentTaxiRequestInfo":
		return DriverCurrentTaxiRequestInfo, nil
	case "driverGetOffers":
		return DriverGetOffers, nil
	}

	return UnknownWsClientType, errors.New("Unknown WsClientType: " + str)
}