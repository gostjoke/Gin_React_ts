package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"gin-backend/config"
	"gin-backend/models"
)

var JWT_SECRET = []byte("MY_SECRET_KEY")

func Register(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 密碼加密
	hash, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	user := models.User{
		Username: body.Username,
		Password: string(hash),
	}

	config.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User registered"})
}

func Login(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user models.User
	config.DB.Where("username = ?", body.Username).First(&user)

	// 密碼比對
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong username or password"})
		return
	}

	// 建立 JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // 一天有效
	})

	tokenString, _ := token.SignedString(JWT_SECRET)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
