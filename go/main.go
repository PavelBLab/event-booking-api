package main

import (
	"github.com/PavelBLab/event-booking-api/configurations/postgres"
	"github.com/PavelBLab/event-booking-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	//sqllite3.InitDB()
	postgres.InitDB()

	server := gin.Default()

	routes.RegisterRoutes(server)

	err := server.Run(":8080") // localhost:8080

	if err != nil {
		return
	}
}
