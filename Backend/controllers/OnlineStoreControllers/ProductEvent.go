package onlinestorecontrollers

type CreateProductRequest struct {
	// ID          uint    
	Name        string  `form:"Name" json:"Name" binding:"required"`
	Description string  `form:"Description" json:"Description"`
	Price      float64  `json:"price" binding:"required,gt=0"`
	// Stock       int     
	// IsActive    bool   
	CategoryID uint     `json:"category_id" binding:"required"`
	// Category   Category `form:"Category" json:"Category"`

	// CreatedAt time.Time
	// UpdatedAt time.Time
}

func CreateProduct(c *gin.Context) {
	var req CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "invalid request data",
			"detail": err.Error(),
		})
		return
	}

	err := config.DB.Transaction(func(tx *gorm.DB) error {

		// 1️⃣ 確認 Category 是否存在（等價 Django .get()）
		var category models.Category
		if err := tx.First(&category, req.CategoryID).Error; err != nil {
			return errors.New("category not found")
		}

		// 2️⃣ 建立 Product
		product := models.Product{
			Name:        req.Name,
			Description: req.Description,
			Price:       req.Price,
			CategoryID:  req.CategoryID,
		}

		if err := tx.Create(&product).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "product created successfully",
	})
}