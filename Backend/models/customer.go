package models

import "time"

// =======================
// Customer（客戶）
// =======================
// SAP: KNA1
//
type Customer struct {
	Code      string `gorm:"primaryKey;size:20"` // APPLE / DELL
	Name      string `gorm:"size:100"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
