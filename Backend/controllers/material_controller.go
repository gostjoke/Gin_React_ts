package controllers

import (
	"fmt"
	"gin-backend/config"
	"gin-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ManufacturerInsert(c *gin.Context) {
	// Post
	Code := c.PostForm("Code")
	Name := c.PostForm("Name")

	if Code == "" || Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code and Name required"})
		return
	}

	Manufacturer := models.Manufacturer{
		Code: Code,
		Name: Name,
	}

	config.DB.Create(&Manufacturer)
	fmt.Println("Registered code: ", Code, ", name: ", Name)
	message := fmt.Sprintf("Code %s registered successfully", Code)
	c.JSON(200, gin.H{"message": message})
}

func GetManufacturerList(c *gin.Context) {
	var manufacturers []models.Manufacturer
	result := config.DB.Limit(100).Find(&manufacturers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve manufacturers"})
		return
	}
	c.JSON(http.StatusOK, manufacturers)
}
