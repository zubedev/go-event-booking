package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-event-booking/db"
	"go-event-booking/routes"
	"net/http"
)

func main() {
	// Initialize database
	db.InitDB()

	// Initialize server
	server := gin.Default()

	// Routes
	server.GET("/", getRoot)
	routes.RegisterRouter(server)

	// Start server
	err := server.Run(":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getRoot(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Event Booking API",
		"version": "0.1.0",
		"author":  "Zubair Beg",
	})
}
