package models

type Category struct {
	No			string 		`gorm:"primaryKey;size:30"`
	Description sting  		`gorm:"size:255"`
	CreatedAt  	time.Time 	`gorm:"autoCreateTime"`
	UpdatedAt  	time.Time 	`gorm:"autoUpdateTime"`
}


type SparePart struct {
	Number     	string 		`gorm:"primaryKey;size:30"`
	Category_no	Category	`gorm:"foreignKey:Category;references:No"`
	CreatedAt  	time.Time	`gorm:"autoCreateTime"`
	UpdatedAt  	time.Time 	`gorm:"autoUpdateTime"`
}
