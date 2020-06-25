package services

type GoogleMapsServiceMock struct {}

var (
	RunGetPlaceDetails func(placeId string) (string, error)
)

func (q *GoogleMapsServiceMock) GetPlaceDetails(placeId string) (string, error) {

	return RunGetPlaceDetails(placeId)
}