package models

import "time"

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"size:200;not null"`
	Description string  `gorm:"type:text"`
	Price       float64 `gorm:"type:decimal(10,2);not null"`
	Stock       int     `gorm:"not null;default:0"`
	IsActive    bool    `gorm:"default:true"`

	CategoryID uint
	Category   Category

	Images []ProductImage

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ProductImage struct {
	ID        uint `gorm:"primaryKey"`
	ProductID uint
	ImageURL  string `gorm:"size:255;not null"`
	IsMain    bool   `gorm:"default:false"`
}

type Category struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:100;uniqueIndex;not null"`

	Products []Product
}

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint
	ProductID uint
	Product   Product

	Qty int `gorm:"not null"`
}

type Cart struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	User   User

	Items []CartItem

	CreatedAt time.Time
	UpdatedAt time.Time
}

type OnlineStoreOrder struct {
	ID uint `gorm:"primaryKey"`

	UserID uint
	User   User

	Status string `gorm:"size:50;default:'pending'"`
	Total  float64

	Items   []OrderItem `gorm:"foreignKey:OrderID"`
	Payment *Payment    `gorm:"foreignKey:OrderID"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type Payment struct {
	ID uint `gorm:"primaryKey"`

	OrderID uint              `gorm:"uniqueIndex"`
	Order   *OnlineStoreOrder `gorm:"constraint:OnDelete:CASCADE"`

	Method string  `gorm:"size:50"`
	Status string  `gorm:"size:50"`
	Amount float64 `gorm:"not null"`

	PaidAt    *time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Product   Product

	Price float64 `gorm:"not null"` // 當下成交價
	Qty   int     `gorm:"not null"`
}
