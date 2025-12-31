package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lawalajose/go-rest-api/models"
)

func main() {
	server := gin.Default()

	event1 := models.Event{
		ID:          1,
		Name:        "Jummah",
		Description: "Congregational Prayer",
		Location:    "Central Mosque",
		UserID:      1,
	}
	event2 := models.Event{
		ID:          2,
		Name:        "Aqdu",
		Description: "Marriage Ceremony",
		Location:    "Mosque Hall",
		UserID:      2,
	}

	event1.Save()
	event2.Save()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	fmt.Println("Server running on port 8080")
	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Message": "Could not parse request data"})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	context.JSON(http.StatusCreated, gin.H{"Message": "event created successfuly", "event": event})

}
