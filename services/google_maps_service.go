package services

import "github.com/testing/external/clients/googlemaps"

type GoogleMapsServiceInterface interface {
	GetPlaceDetails(string) (string, error)
}

type googleMapsService struct {}

var (
	GoogleMapsService GoogleMapsServiceInterface = &googleMapsService{}
	GoogleMapsClient  googlemaps.GoogleMapsClientInterface
)

func init() {
	GoogleMapsClient = googlemaps.GoogleMapsClient
}

func (g *googleMapsService) GetPlaceDetails(placeId string) (string, error) {

	return GoogleMapsClient.Get(placeId)
}