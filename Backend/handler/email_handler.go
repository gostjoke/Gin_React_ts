package handler

import (
	"gin-backend/domain"
	"gin-backend/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailHandler struct {
	Service service.EmailService
}

type EmailRequest struct {
	To      []string `json:"to" binding:"required"`
	CC      []string `json:"cc"`
	Subject string   `json:"subject" binding:"required"`
	Body    string   `json:"body" binding:"required"`
}

func (h *EmailHandler) Send(c *gin.Context) {
	var req EmailRequest

	// 1️⃣ Bind & validate request
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 2️⃣ Map request → domain
	email := domain.Email{
		From:    "your@gmail.com",
		To:      req.To,
		CC:      req.CC,
		Subject: req.Subject,
		Body:    req.Body,
	}

	// 3️⃣ Call service
	if err := h.Service.Send(email); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "failed to send email",
			"detail": err.Error(),
		})
		return
	}

	// 4️⃣ Response
	c.JSON(http.StatusOK, gin.H{
		"message": "email sent successfully",
	})
}
