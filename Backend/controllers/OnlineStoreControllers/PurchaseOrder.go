package onlinestorecontrollers

import (
	"gin-backend/config"
	"gin-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePurchaseOrder(c *gin.Context) error {
	// TODO: Implement the logic to create a purchase order
	Code := c.PostForm("Code")
	Name := c.PostForm("Name")
	ProductID := c.PostForm("ProductID")

	if Code == "" || Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code and Name required"})
		return
	}

	// atomic transaction 原子操作
	err := config.DB.Transaction(func(tx *gorm.DB) error {
		OrderItem := models.OrderItem{

			Product: models.Product{
				ID:          ProductID,
				Name:        "",
				Description: "",
				Price:       0,
				Stock:       0,
				IsActive:    true,
				CategoryID:  0,
			},

			Price: 0, // 當下成交價
			Qty:   0,
		}
	})

	// 1️⃣ 插入資料
	if err := tx.Create(&OrderItem).Error; err != nil {
		return err // ❌ 自動 rollback
	}

	return nil
}
