package infrastructure

import (
	"fmt"
	"gin-backend/domain"
	"net/smtp"
	"strings"
)

type GmailService struct {
	// Add necessary fields for Gmail service configuration
	Email    string
	Password string
}

func (g *GmailService) Send(email domain.Email) error {
	host := "smtp.gmail.com"
	port := "587"
	addr := host + ":" + port

	// Gmail SMTP auth
	auth := smtp.PlainAuth(
		"",
		g.Email,
		g.Password, // App Password
		host,
	)

	// ===== Build headers =====
	headers := map[string]string{
		"From":         email.From,
		"To":           strings.Join(email.To, ", "),
		"Subject":      email.Subject,
		"MIME-Version": "1.0",
		"Content-Type": `text/html; charset="UTF-8"`,
	}

	if len(email.CC) > 0 {
		headers["Cc"] = strings.Join(email.CC, ", ")
	}

	// ===== Combine headers =====
	var msg strings.Builder
	for k, v := range headers {
		msg.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	msg.WriteString("\r\n")     // Header / body separator
	msg.WriteString(email.Body) // HTML body

	// ===== All recipients =====
	recipients := append(email.To, email.CC...)

	// ===== Send =====
	return smtp.SendMail(
		addr,
		auth,
		email.From,
		recipients,
		[]byte(msg.String()),
	)
}
