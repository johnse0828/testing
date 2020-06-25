package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/testing/services"
	"net/http/httptest"
	"testing"
)

func init() {
	services.GoogleMapsService = &services.GoogleMapsServiceMock{}
}

func TestGetPlaceDetails(t *testing.T) {
	services.RunGetPlaceDetails = func(string) (string, error) {
		return "", nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Query("")
	GetPlaceDetails(c)
	assert.Equal(t, 200, w.Code)
}
