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

func main() {
	r := gin.Default()

	// Connect DB
	config.ConnectDB()
	config.DB.AutoMigrate(&models.Department{})
	config.DB.AutoMigrate(&models.UserProfile{})
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Customer{})
	config.DB.AutoMigrate(&models.SerialNumber{})
	config.DB.AutoMigrate(&models.CPN{})
	config.DB.AutoMigrate(&models.Rma{})

	// Routes
	routes.RegisterRoutes(r)

	cwd, _ := os.Getwd()
	log.Println("Current working dir:", cwd)

	// Email
	if err := godotenv.Load(); err != nil {
		fmt.Println("⚠️ No .env file found, using system env")
		err := os.WriteFile(".env.example", []byte("GMAIL_USER=\nGMAIL_PASS=\n"), 0644)
		if err != nil {
			fmt.Println("⚠️ No .env.example file found and failed to create one:", err)
		}
	}

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
