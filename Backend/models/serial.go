package models

type SerialNumber struct {
	Number     string `gorm:"primaryKey;size:30"`
	CPNID      uint
	CPN        CPN  `gorm:"foreignKey:CPNID;references:ID"`
	IsAssigned bool `gorm:"default:false"`
}
