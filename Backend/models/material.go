package models

import "time"

//
// =======================
// Manufacturer（製造商）
// =======================
// SAP: LFA1 / Manufacturer
//
type Manufacturer struct {
	Code      string `gorm:"primaryKey;size:20"` // TI / ADI / NXP
	Name      string `gorm:"size:100"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//
// =======================
// MPN（Manufacturer Part Number）
// =======================
// SAP: MFR + MFRPN
// Composite PK: Manufacturer + MPN
//
type MPN struct {
	ManufacturerCode string       `gorm:"primaryKey;size:20"`
	Number           string       `gorm:"primaryKey;size:50"`
	Manufacturer     Manufacturer `gorm:"foreignKey:ManufacturerCode;references:Code"`

	Description string `gorm:"size:255"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//
// =======================
// MATNR（SAP Material Master）
// =======================
// SAP: MARA / MARC
//
type MATNR struct {
	Number       string `gorm:"primaryKey;size:40"` // SAP MATNR
	Description  string `gorm:"size:255"`
	MaterialType string `gorm:"size:10"` // ROH / HALB / FERT
	BaseUnit     string `gorm:"size:10"` // EA / KG

	CreatedAt time.Time
	UpdatedAt time.Time

	MPNLinks []MATNRMPN `gorm:"foreignKey:MATNRNumber;references:Number"`
	CPNLinks []CPN      `gorm:"foreignKey:MATNRNumber;references:Number"`
}

//
// =======================
// MATNR ↔ MPN（AVL / Second Source）
// =======================
// SAP: Source List / AVL 強化版
//
type MATNRMPN struct {
	MATNRNumber      string `gorm:"primaryKey;size:40"`
	ManufacturerCode string `gorm:"primaryKey;size:20"`
	MPNNumber        string `gorm:"primaryKey;size:50"`

	MATNR MATNR `gorm:"foreignKey:MATNRNumber;references:Number"`
	MPN   MPN   `gorm:"foreignKey:ManufacturerCode,MPNNumber;references:ManufacturerCode,Number"`

	Status    string `gorm:"size:20"` // Approved / Blocked
	IsPrimary bool
	CreatedAt time.Time
}

//

//
// =======================
// CPN（Customer Part Number）
// =======================
// SAP: KNMT（Customer-Material Info Record）
// Composite PK: Customer + CPN + Revision
//
type CPN struct {
	CustomerCode string `gorm:"primaryKey;size:20"`
	Number       string `gorm:"primaryKey;size:50"` // CPN
	Revision     string `gorm:"primaryKey;size:10"`

	MATNRNumber string `gorm:"size:40"`
	MATNR       MATNR  `gorm:"foreignKey:MATNRNumber;references:Number"`

	Customer Customer `gorm:"foreignKey:CustomerCode;references:Code"`

	Active    bool
	CreatedAt time.Time
}
