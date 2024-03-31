package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-event-booking/utils"
	"net/http"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil || userId == 0 {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
