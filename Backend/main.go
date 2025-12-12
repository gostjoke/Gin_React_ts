package main

import (
	"gin-backend/config"
	"gin-backend/models"
	"gin-backend/routes"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {
	r := gin.Default()

	// Connect DB
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Department{})
	config.DB.AutoMigrate(&models.UserProfile{})
	config.DB.AutoMigrate(&models.User{})

	// Routes
	routes.RegisterRoutes(r)

	// Run server
	r.Run(":8080")
}
