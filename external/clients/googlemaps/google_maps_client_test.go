package googlemaps

import (
	"bytes"
	"fmt"
	"github.com/testing/utils/mocks"
	"io/ioutil"
	"net/http"
	"testing"
)

func init() {
	Client = &mocks.MockClient{}
}

func TestGetPlaceDetails(t *testing.T) {
	placeId := "ChIJN1t_tDeuEmsRUsoyG83frY4"

	responseMocked := ioutil.NopCloser(bytes.NewReader([]byte("my response")))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body: responseMocked,
		}, nil
	}
	resp, err := GoogleMapsClient.Get(placeId)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
}