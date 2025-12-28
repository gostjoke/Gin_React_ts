package routes

import (
	"gin-backend/controllers"

	"github.com/gin-gonic/gin"
)

func MaterialRoutes(r *gin.Engine) {
	r.POST("/ManufacturerInsert", controllers.ManufacturerInsert)
	r.GET("/GetManufacturerList", controllers.GetManufacturerList)
}
