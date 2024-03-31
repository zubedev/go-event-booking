package routes

import (
	"github.com/gin-gonic/gin"
	"go-event-booking/models"
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
