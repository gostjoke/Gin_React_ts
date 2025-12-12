package models

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	Name      string `json:"name" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Department) TableName() string {
	return "django_departments"
}

type User struct {
	Username  string `json:"username" gorm:"primaryKey"`
	Password  string `json:"-" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// 一對一
	Profile UserProfile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserProfile struct {
	gorm.Model
	UserID         uint       `json:"user_id" gorm:"uniqueIndex"`
	FirstName      string     `json:"first_name"`
	LastName       string     `json:"last_name"`
	Email          string     `json:"email" gorm:"unique"`
	Phone          string     `json:"phone"`
	Address        string     `json:"address"`
	DepartmentName string     `json:"department_name"`
	Department     Department `gorm:"foreignKey:DepartmentName;references:Name"`
}
