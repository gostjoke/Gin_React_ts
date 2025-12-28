package models

import (
	"time"

	"gorm.io/gorm"
)

/* ---------------- Department ---------------- */

type Department struct {
	Name      string `gorm:"primaryKey;size:191"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Department) TableName() string {
	return "django_departments"
}

/* ---------------- User ---------------- */

type User struct {
	Username  string `json:"username" gorm:"primaryKey;size:191"`
	Password  string `json:"-" gorm:"not null"` // Hashed password and not return serialized
	CreatedAt time.Time
	UpdatedAt time.Time
}

/* ---------------- UserProfile ---------------- */

type UserProfile struct {
	Username string `gorm:"primaryKey;size:191"`

	// belongs to User
	User User `gorm:"foreignKey:Username;references:Username;constraint:OnDelete:CASCADE;"`

	FirstName string
	LastName  string
	Email     string `gorm:"size:191;unique"`
	Phone     string
	Address   string

	DepartmentName string     `gorm:"size:191;index"`
	Department     Department `gorm:"foreignKey:DepartmentName;references:Name"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
