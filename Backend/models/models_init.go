package models

import (
	"gin-backend/config"
)

func InitOnlineStoreModels() {
	// Connect DB
	// config.ConnectDB()
	config.ConnectPostgreDB()

	config.DB.AutoMigrate(&Department{})
	config.DB.AutoMigrate(&User{})
	config.DB.AutoMigrate(&UserProfile{})

	config.DB.AutoMigrate(&Customer{})

	config.DB.AutoMigrate(&Manufacturer{})
	config.DB.AutoMigrate(&MATNR{})
	config.DB.AutoMigrate(&MPN{})
	config.DB.AutoMigrate(&MATNRMPN{})
	config.DB.AutoMigrate(&CPN{})
	config.DB.AutoMigrate(&Rma{})
	config.DB.AutoMigrate(&SerialNumber{})

	// Online Store Models
	config.DB.AutoMigrate(&Category{})
	config.DB.AutoMigrate(&Product{})
	config.DB.AutoMigrate(&ProductImage{})

	config.DB.AutoMigrate(&OrderItem{})
	config.DB.AutoMigrate(&OnlineStoreOrder{})
	config.DB.AutoMigrate(&Payment{})

	config.DB.AutoMigrate(&Cart{})
	config.DB.AutoMigrate(&CartItem{})
}
