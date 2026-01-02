package models

import "time"

type SerialNumber struct {
	Number     string `gorm:"primaryKey;size:30"`
	CPNID      uint
	CPN        CPN       `gorm:"foreignKey:CPNID;references:ID"`
	IsAssigned bool      `gorm:"default:false"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
