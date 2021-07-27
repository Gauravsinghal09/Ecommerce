package Product

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ProductName     string
	ProductPrice    int
	ProductQuantity int `gorm:"check:quantity >= 0"`
}

func (product *Products) TableName() string {
	return "products"
}
