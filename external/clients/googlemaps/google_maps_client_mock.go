package googlemaps

type GoogleMapsClientMock struct {
}

var (
	RunGetFunc func(placeId string) (string, error)
)

func (g *GoogleMapsClientMock) Get(placeId string) (string, error) {
	return RunGetFunc(placeId)
}