package main

import (
	"fmt"
	"gin-backend/handler"
	"gin-backend/infrastructure"
	"gin-backend/models"
	"gin-backend/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func init() {
	// Load .env file before any other package initialization
	if err := godotenv.Load(); err != nil {
		fmt.Println("⚠️ No .env file found, using system env")
	}
}

func main() {
	r := gin.Default()

	// Initialize DB and Models
	models.InitOnlineStoreModels()

	// Test Routes, must close before production danger!!!!
	routes.TestRoutes(r)

	// Routes
	routes.RegisterRoutes(r)
	routes.MaterialRoutes(r)

	cwd, _ := os.Getwd()
	log.Println("Current working dir:", cwd)

	// Email setup
	gmailUser := os.Getenv("GMAIL_USER")
	gmailPass := os.Getenv("GMAIL_PASS")

	gmailService := &infrastructure.GmailService{
		Email:    gmailUser,
		Password: gmailPass,
	}

	fmt.Println("GMAIL_USER: ", gmailUser)
	fmt.Println("GMAIL_PASS: ", gmailPass)

	emailHandler := &handler.EmailHandler{
		Service: gmailService, // 👈 自動符合 EmailInterface
	}

	r.POST("/email/send", emailHandler.Send)

	// Run server
	r.Run(":8080")
}
