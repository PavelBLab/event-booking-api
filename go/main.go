package main

import (
	"net/http"

	"github.com/PavelBLab/event-booking-api/configurations/postgres"
	"github.com/PavelBLab/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func main() {
	//sqllite3.InitDB()
	postgres.InitDB()

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":8080") // localhost:8080

	if err != nil {
		return
	}
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve events: " + err.Error()})
		return
	}

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request:  " + err.Error()})
		return
	}

	event.ID = 1
	event.UserId = 1
	err = event.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event: " + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, event)

}
