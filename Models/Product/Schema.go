package Product

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	ProductName     string
	ProductPrice    int
	ProductQuantity int
}

func (product *Products) TableName() string {
	return "products"
}
