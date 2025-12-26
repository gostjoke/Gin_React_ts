package service

import "gin-backend/domain"

type EmailService interface {
	Send(email domain.Email) error
}
