package models

import "gorm.io/gorm"

type Department struct {
	gorm.Model
	Name string `json:"name" gorm:"unique"`
}

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
}

type UserProfile struct {
	gorm.Model
	UserID       uint   `json:"user_id" gorm:"uniqueIndex"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" gorm:"unique"`
	Phone        string `json:"phone" gorm:"unique"`
	Address      string `json:"address"`
	DepartmentID uint   `json:"department_id"`
	Department
}
