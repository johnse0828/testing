package main

import (
	"github.com/gin-gonic/gin"
	"github.com/testing/controllers"
	"log"
)

func helloWorld(c *gin.Context) {
	c.JSON(200, "Hello World!")
}

func main() {
	r := gin.Default()

	r.GET("/", helloWorld)
	r.GET("/place-details/", controllers.GetPlaceDetails)

	log.Fatal(r.Run())
}