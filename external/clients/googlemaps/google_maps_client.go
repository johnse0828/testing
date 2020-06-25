package googlemaps

import (
	"io/ioutil"
	"net/http"
)

var (
	GoogleMapsClient GoogleMapsClientInterface = &googleMapsClient{}
	url                                        = "https://maps.googleapis.com/maps/api/place/details/json?key=AIzaSyAPf4zyVQmEK7l8WGYsehnstxg7aOdV_Cw&place_id="
	Client           HTTPClient
)

type googleMapsClient struct {}

type GoogleMapsClientInterface interface {
	Get(placeId string) (string, error)
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type GoogleErr struct {
	error `json:"error"`
}

func init() {
	Client = &http.Client{}
}

func (g *googleMapsClient) Get(placeId string) (string, error) {
	request, err := http.NewRequest("GET", url+ placeId, nil)

	if err != nil {
		return "", GoogleErr{err }
	}

	response, err := Client.Do(request)

	if err != nil {
		return "", GoogleErr{err }
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", GoogleErr{err }
	}

	respAsString := string(body)

	return respAsString, nil
}