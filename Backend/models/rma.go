package models

import "time"

type Rma struct {
	Number       string       `gorm:"primaryKey;size:20"` // RMA1234567890
	SerialNumber SerialNumber `gorm:"foreignKey:SerialNumber;references:Number"`
	Status       string       `gorm:"size:20"` // Pending / Approved / Rejected / Completed
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
