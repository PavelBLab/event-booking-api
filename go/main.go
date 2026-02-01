package main

import (
	"net/http"

	"github.com/PavelBLab/event-booking-api/models"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	err := server.Run(":8080") // localhost:8080

	if err != nil {
		return
	}
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	event.ID = 1
	event.UserId = 1
	event.Save()

	context.JSON(http.StatusCreated, event)

}
