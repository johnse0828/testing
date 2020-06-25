package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/testing/services"
	"log"
)

func GetPlaceDetails(c *gin.Context) {
	placeId := c.Query("place_id")
	resp, err := services.GoogleMapsService.GetPlaceDetails(placeId)

	if err != nil {
		log.Fatal(err)
	}

	c.Data(200, "application/json", []byte(resp))
}