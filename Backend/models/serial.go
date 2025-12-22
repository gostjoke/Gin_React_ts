package models

type SerialNumber struct {
	Number     string `gorm:"primaryKey;size:30"` // SN1234567890
	CPNNumber  string `gorm:"size:50"`
	CPN        CPN    `gorm:"foreignKey:CPNNumber;references:Number"`
	IsAssigned bool   `gorm:"default:false"`
}
