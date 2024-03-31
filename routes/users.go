package routes

import (
	"github.com/gin-gonic/gin"
	"go-event-booking/models"
	"go-event-booking/utils"
	"net/http"
)

func signup(context *gin.Context) {
	var user *models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	user, err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	context.JSON(http.StatusCreated, user)
}

func login(context *gin.Context) {
	var user *models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	user, err = user.ValidateCredentials()
	if err != nil || user == nil || user.ID == 0 {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil || token == "" {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
