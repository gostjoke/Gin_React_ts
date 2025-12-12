package models

import (
	"time"

	"gorm.io/gorm"
)

type SafetyStockControl struct {
	gorm.Model
	ItemCode     string    `json:"item_code" gorm:"uniqueIndex"`
	ItemName     string    `json:"item_name"`
	CurrentStock int       `json:"current_stock"`
	SafetyStock  int       `json:"safety_stock"`
	ReorderPoint int       `json:"reorder_point"`
	LeadTimeDays int       `json:"lead_time_days"`
	DailyUsage   int       `json:"daily_usage"`
	LastUpdated  time.Time `json:"last_updated"`
}

func (SafetyStockControl) TableName() string {
	return "safety_stock_controls"
}
