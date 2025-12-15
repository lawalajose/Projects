package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/db"
	"go-rest-api/models"
	"net/http"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents) //GET, POST, PUSH, PATCH, DELETE
	server.POST("/events", createEvent)

	server.Run(":8080") //Localhost

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})

}
