package onlinestorecontrollers

import (
	"gin-backend/config"
	"gin-backend/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateCategoryRequest struct {
	Name        string `form:"Name" json:"Name" binding:"required"`
	Description string `form:"Description" json:"Description" binding:"required"`
}

func CreateCategory(c *gin.Context) {
	var req CreateCategoryRequest

	// 解析 POST form 或 JSON
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Missing required fields",
			"detail": err.Error(),
		})
		return
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {

		category := models.Category{
			Name:        req.Name,
			Description: req.Description,
		}

		if err := tx.Create(&category).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category created successfully",
	})
}

func CreatePurchaseOrder(c *gin.Context) {
	// TODO: Implement the logic to create a purchase order
	Code := c.PostForm("Code")
	Name := c.PostForm("Name")
	ProductIDStr := c.PostForm("ProductID")

	if strings.TrimSpace(Code) == "" || strings.TrimSpace(Name) == "" || strings.TrimSpace(ProductIDStr) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Code, Name, ProductID are required"})
		return
	}

	// ProductID 轉換成 uint
	productID, err := strconv.ParseUint(ProductIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ProductID"})
		return
	}

	err = config.DB.Transaction(func(tx *gorm.DB) error {

		orderItem := models.OrderItem{
			ProductID: uint(productID), // 外鍵
			Price:     0,
			Qty:       0,
		}

		// 插入資料
		if err := tx.Create(&orderItem).Error; err != nil {
			return err // rollback
		}

		return nil // commit
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}


