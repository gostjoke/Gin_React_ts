package main

import (
	"fmt"
	"gin-backend/config"
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

	// Connect DB
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Department{})
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.UserProfile{})

	config.DB.AutoMigrate(&models.Customer{})

	config.DB.AutoMigrate(&models.Manufacturer{})
	config.DB.AutoMigrate(&models.MATNR{})
	config.DB.AutoMigrate(&models.MPN{})
	config.DB.AutoMigrate(&models.MATNRMPN{})
	config.DB.AutoMigrate(&models.CPN{})
	config.DB.AutoMigrate(&models.Rma{})
	config.DB.AutoMigrate(&models.SerialNumber{})

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
