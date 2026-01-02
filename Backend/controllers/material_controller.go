package controllers

import (
	"fmt"
	"gin-backend/config"
	"gin-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ManufacturerInsert(c *gin.Context) {
	Code := c.PostForm("Code")
	Name := c.PostForm("Name")

	if Code == "" || Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code and Name required"})
		return
	}

	// atomic transaction 原子操作
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		manufacturer := models.Manufacturer{
			Code: Code,
			Name: Name,
		}

		// 1️⃣ 插入資料
		if err := tx.Create(&manufacturer).Error; err != nil {
			return err // ❌ 自動 rollback
		}

		// 2️⃣ 你未來可以加其他 DB 操作
		// if err := tx.Create(&log).Error; err != nil {
		//     return err
		// }

		return nil // ✅ commit
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Insert failed",
			"detail": err.Error(),
		})
		return
	}

	fmt.Println("Registered code:", Code, ", name:", Name)
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Code %s registered successfully", Code),
	})
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
