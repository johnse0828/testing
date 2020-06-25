package services

import (
	"fmt"
	"github.com/testing/external/clients/googlemaps"
	"testing"
)

func init() {
	GoogleMapsClient = &googlemaps.GoogleMapsClientMock{}
}

func TestGoogleMapsService_GetPlaceDetails(t *testing.T) {
	placeId := "ChIJN1t_tDeuEmsRUsoyG83frY4"
	googlemaps.RunGetFunc = func(string) (string, error) {
		return "Hello", nil
	}

	resp, err := GoogleMapsService.GetPlaceDetails(placeId)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}