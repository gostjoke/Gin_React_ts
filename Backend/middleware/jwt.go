package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var JWT_SECRET []byte

func init() {
	// Load .env file if it exists
	_ = godotenv.Load()

	JWT_SECRET = []byte(os.Getenv("JWT_SECRET"))
	if len(JWT_SECRET) == 0 {
		panic("JWT_SECRET is not set")
	}
	fmt.Println("✅ JWT_SECRET loaded successfully")
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 取得 Authorization: Bearer xxx
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return JWT_SECRET, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
