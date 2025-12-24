package models

import "time"

type Rma struct {
	Number       string       `gorm:"primaryKey;size:20"`
	SerialNumber SerialNumber `gorm:"foreignKey:Number;references:Number"`

	Status    string `gorm:"size:20"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
