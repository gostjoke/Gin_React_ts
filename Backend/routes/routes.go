package routes

import (
	"github.com/gin-gonic/gin"

	"gin-backend/controllers"
	"gin-backend/middleware"
)

func RegisterRoutes(r *gin.Engine) {
	r.Use(middleware.CORS())
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/metrics/qps", func(c *gin.Context) {
		c.JSON(200, gin.H{"qps": middleware.GetQPS()})
	})

	// 受保護的 API
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "You are authorized"})
		})
	}
}

func TestRoutes(r *gin.Engine) {
	r.POST("/test/user", controllers.GetUserList)
}
