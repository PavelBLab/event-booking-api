package routes

import (
	"net/http"

	"github.com/PavelBLab/event-booking-api/models"
	"github.com/PavelBLab/event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request:  " + err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user: " + err.Error()})
		return
	}

	user.Password = ""
	context.JSON(http.StatusCreated, user)
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Bad request:  " + err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate JWT: " + err.Error()})
		return
	}

	user.Password = ""
	context.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}
