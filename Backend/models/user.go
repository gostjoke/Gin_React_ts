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

	// One-to-One relationship with UserProfile
	Profile *UserProfile `gorm:"foreignKey:Username;references:Username;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type UserProfile struct {
	Username string `json:"username" gorm:"primaryKey"`
	User     User   `gorm:"foreignKey:Username;references:Username"`

	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Phone     string
	Address   string

	DepartmentName string
	Department     Department `gorm:"foreignKey:DepartmentName;references:Name"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
