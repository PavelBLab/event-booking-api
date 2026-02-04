package middlewares

import (
	"net/http"

	"github.com/PavelBLab/event-booking-api/utils"
	"github.com/gin-gonic/gin"
)

func Auth(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyJWT(token)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	context.Set("userId", userId)

	context.Next()
}
